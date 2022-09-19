package mongodb

import "github.com/Ioloman/migration-script/app/models"

func InsertLogs(logs *[]models.PaymentLog) error {
	logsVal := *logs
	insertLogs := make([]interface{}, len(logsVal))
	for i := range logsVal {
		insertLogs[i] = logsVal[i]
	}

	_, err := Collection.InsertMany(CTX, insertLogs)
	return err
}
