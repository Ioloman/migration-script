package parallel

import (
	"log"
	"time"

	"github.com/Ioloman/migration-script/app/db/mysql"
	"github.com/Ioloman/migration-script/app/db/mongodb"
	"github.com/Ioloman/migration-script/app/models"
)

func worker(inputCh <-chan *[]uint64, outputCh chan<- models.WorkerReturn) {
	for {
		localTiming := &models.Timings{Count: 1}
		result := models.WorkerReturn{Stats: localTiming}

		paymentIDs := <-inputCh
		result.PaymentIDs = paymentIDs
		localTiming.LogsAmount = uint64(len(*paymentIDs))
		
		t := time.Now()

		logs, err := mysql.GetLogs(paymentIDs)
		if err != nil {
			result.Error = err
			outputCh <- result
			continue
		}
		if len(*logs) == 0 {
			continue
		}
		t = localTiming.SetSelect(t)

		err = mongodb.InsertLogs(logs)
		if err != nil {
			result.Error = err
			outputCh <- result
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

		outputCh <- result
	}
}

func Migrate(batchSize int, numWorkers int, printEvery int) error {
	var returnBuffer models.WorkerReturn
	globalTiming := &models.Timings{}
	inputCh := make(chan *[]uint64, numWorkers)
	outputCh := make(chan models.WorkerReturn, numWorkers)
	paymentID := uint64(0)

	for i := 0; i < numWorkers; i++ {
		go worker(inputCh, outputCh)
	}

	for {
		t := time.Now()
		paymentIDs, err := mysql.GetPaymentIDs(batchSize, paymentID)
		if err != nil {
			log.Fatalf("Error querying payment_ids: %v", err)
			continue
		}
		if len(*paymentIDs) == 0 {
			log.Println("Got 0 logs")
			time.Sleep(time.Second * 5)
			continue
		}
		globalTiming.AddSelect(t)
		paymentID = (*paymentIDs)[len(*paymentIDs)-1]

		select {
		case inputCh <- paymentIDs:

		case returnBuffer = <-outputCh:
			if returnBuffer.Error != nil {
				log.Fatalf("got output with error: %v\n", returnBuffer.Error)
				inputCh <- returnBuffer.PaymentIDs
			} else {
				globalTiming.Add(returnBuffer.Stats)
			}
		}

		if globalTiming.Count%uint64(printEvery) == 0 {
			log.Printf("globalTiming: %v\n", globalTiming)
		}
	}
}
