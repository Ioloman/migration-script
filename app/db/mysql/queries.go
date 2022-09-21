package mysql

import (
	"fmt"

	"github.com/Ioloman/migration-script/app/models"
	"github.com/jmoiron/sqlx"
)

func DeleteLogs(paymentIDs *[]uint64, database string) error {
	query, args, err := sqlx.In(fmt.Sprintf("DELETE FROM %v WHERE payment_id in (?)", database), *paymentIDs)
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

func GetPaymentIDs(n int, database string) (*[]uint64, error) {
	paymentIDs := []uint64{}

	err := DB.Select(
		&paymentIDs,
		fmt.Sprintf("SELECT DISTINCT payment_id FROM %v WHERE payment_id IS NOT NULL AND payment_id > 0 ORDER BY payment_id ASC LIMIT ?", database),
		n,
	)
	if err != nil {
		return nil, err
	}

	return &paymentIDs, nil
}

func GetLogs(paymentIDs *[]uint64, database string) (*[]models.PaymentLog, error) {
	logs := []models.PaymentLog{}

	query, args, err := sqlx.In(fmt.Sprintf("SELECT payment_id, text, date FROM %v WHERE payment_id IN (?)", database), *paymentIDs)
	if err != nil {
		return nil, err
	}
	query = DB.Rebind(query)
	err = DB.Select(&logs, query, args...)
	return &logs, err
}
