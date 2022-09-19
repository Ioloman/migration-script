package single

import (
	"log"
	"time"

	"github.com/Ioloman/migration-script/app/db/mongodb"
	"github.com/Ioloman/migration-script/app/db/mysql"
	"github.com/Ioloman/migration-script/app/models"
)

func Migrate(batchSize int, printEvery int) error {
	log.Println("starting single migration")
	globalTiming := &models.Timings{NumWorkers: 1}
	paymentID := uint64(0)
	for {
		localTiming := &models.Timings{Count: 1}
		t := time.Now()

		paymentIDs, err := mysql.GetPaymentIDs(batchSize, paymentID)
		if err != nil {
			log.Fatalf("Error querying payment_ids: %v", err)
			continue
		}
		localTiming.LogsAmount = uint64(len(*paymentIDs))
		logs, err := mysql.GetLogs(paymentIDs)
		if err != nil || len(*logs) == 0 {
			log.Fatalf("error or 0 payment_ids: %v", err)
			continue
		}
		if len(*logs) == 0 {
			log.Println("0 logs")
			time.Sleep(time.Second * 5)
			continue
		}
		t = localTiming.SetSelect(t)

		err = mongodb.InsertLogs(logs)
		if err != nil {
			log.Fatal(err)
			continue
		}
		t = localTiming.SetInsert(t)

		err = mysql.DeleteLogs(paymentIDs)
		if err != nil {
			for ; err != nil; err = mysql.DeleteLogs(paymentIDs) {
				log.Fatalf("Cannot delete logs: %v", err)
			}
		}
		localTiming.SetDelete(t)

		globalTiming.Add(localTiming)
		if globalTiming.Count%uint64(printEvery) == 0 {
			log.Printf("localTiming: %v\n", localTiming)
			log.Printf("globalTiming: %v\n", globalTiming)
		}
	}
}
