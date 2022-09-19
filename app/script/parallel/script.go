package parallel

import (
	"github.com/Ioloman/migration-script/app/db/mysql"
	"github.com/Ioloman/migration-script/app/models"
)

func worker(inputCh <-chan *[]uint64, outputCh chan<- *[]uint64) {

}

func Migrate(batchSize int, numWorkers int, printEvery int) error {
	var failedPayments *[]uint64
	globalTiming := &models.Timings{}
	inputCh := make(chan *[]uint64, numWorkers)
	outputCh := make(chan *[]uint64, numWorkers)
	paymentID := uint64(0)

	for i := 0; i < numWorkers; i++ {
		go worker(inputCh, outputCh)
	}

	for {
		paymentIDs := mysql.GetPaymentIDs(batchSize, paymentID)
		paymentID = (*paymentIDs)[len(*paymentIDs)-1]

		select {
		case inputCh <- paymentIDs:

		case failedPayments = <-outputCh:
			inputCh <- failedPayments
		}
	}
}
