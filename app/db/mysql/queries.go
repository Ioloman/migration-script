package mysql

import (
	"github.com/Ioloman/migration-script/app/models"
	"github.com/jmoiron/sqlx"
)

func GetFirstLogs(n int) (*[]models.PaymentLog, error) {
	logs := []models.PaymentLog{}

	err := DB.Select(&logs, "SELECT payment_id, text, date FROM processing.payment_log LIMIT ?", n)
	return &logs, err
}

func DeleteLogs(logs *[]models.PaymentLog) error {
	logIDs := make([]uint64, len(*logs))
	for i, log := range *logs {
		logIDs[i] = log.PaymentId
	}

	query, args, err := sqlx.In("DELETE FROM processing.payment_log WHERE payment_id in (?)", logIDs)
	if err != nil {
		return err
	}
	query = DB.Rebind(query)
	_, err = DB.Exec(query, args...)
	if err != nil {
		return err
	}
	return nil
}
