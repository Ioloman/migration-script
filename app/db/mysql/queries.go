package mysql

import (
	"github.com/Ioloman/migration-script/app/models"
	"github.com/jmoiron/sqlx"
)

func GetFirstLogs(n int) (*[]models.PaymentLog, error) {
	logs := []models.PaymentLog{}

	err := DB.Select(&logs, "SELECT payment_id, text, date FROM processing.payment_log WHERE payment_id IS NOT NULL AND payment_id != 0 LIMIT ?", n)
	return &logs, err
}

func DeleteLogs(paymentIDs *[]uint64) error {
	query, args, err := sqlx.In("DELETE FROM processing.payment_log WHERE payment_id in (?)", *paymentIDs)
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

func GetPaymentIDs(n int, id uint64) (*[]uint64, error) {
	paymentIDs := []uint64{}

	err := DB.Select(
		&paymentIDs,
		"SELECT DISTINCT payment_id FROM processing.payment_log WHERE payment_id IS NOT NULL AND payment_id > ? LIMIT ?",
		id, n,
	)
	if err != nil {
		return nil, err
	}

	return &paymentIDs, nil
}

func GetLogs(paymentIDs *[]uint64) (*[]models.PaymentLog, error) {
	logs := []models.PaymentLog{}

	query, args, err := sqlx.In("SELECT payment_id, text, date FROM processing.payment_log WHERE payment_id IN (?)", *paymentIDs)
	if err != nil {
		return nil, err
	}
	query = DB.Rebind(query)
	err = DB.Select(&logs, query, args...)
	return &logs, err
}
