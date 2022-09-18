package mongodb

import "github.com/Ioloman/migration-script/app/models"

func InsertLogs(logs *[]models.PaymentLog) error {
	insertLogs := make([]interface{}, len(*logs))
	for i, log := range *logs {
		insertLogs[i] = log
	}

	_, err := Collection.InsertMany(CTX, insertLogs)
	return err
}
