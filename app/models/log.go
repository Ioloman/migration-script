package models

import "time"

type PaymentLog struct {
	PaymentId uint32    `db:"payment_id" bson:"payment_id"`
	Text      string    `db:"text" bson:"text"`
	Date      time.Time `db:"date" bson:"date"`
}
