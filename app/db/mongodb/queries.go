package mongodb

import (
	"github.com/Ioloman/migration-script/app/models"
	"go.mongodb.org/mongo-driver/bson"
)

func InsertLogs(logs *[]models.PaymentLog) error {
	logsVal := *logs
	insertLogs := make([]interface{}, len(logsVal))
	for i, log := range logsVal {
		insertLogs[i] = bson.M{"payment_id": log.PaymentId, "date": log.Date, "text": log.Text}
	}

	_, err := Collection.InsertMany(CTX, insertLogs)
	return err
}
