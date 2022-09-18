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
	globalTiming := &models.Timings{}
	for {
		localTiming := &models.Timings{}
		t := time.Now()

		logs, err := mysql.GetFirstLogs(batchSize)
		if err != nil {
			log.Fatal(err)
			continue
		}
		log.Println("got logs")
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
		log.Println("inserted logs")
		t = localTiming.SetInsert(t)

		err = mysql.DeleteLogs(logs)
		if err != nil {
			for ; err != nil; err = mysql.DeleteLogs(logs) {
				log.Fatalf("Cannot delete logs: %v", err)
			}
		}
		log.Println("deleted logs")
		t = localTiming.SetDelete(t)

		globalTiming.Add(localTiming)
		if globalTiming.Count%uint64(printEvery) == 0 {
			log.Printf("localTiming: %v\n", localTiming)
			log.Printf("globalTiming: %v\n", globalTiming)
		}
	}
}
