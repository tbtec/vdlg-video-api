package model

import "time"

type Payment struct {
	ID         string    `gorm:"column:payment_id;primaryKey"`
	OrderId    string    `gorm:"column:order_id"`
	Status     string    `gorm:"column:status"`
	QrData     string    `gorm:"column:qr_data"`
	ExternalId string    `gorm:"column:external_id"`
	CreatedAt  time.Time `gorm:"column:created_at"`
	UpdatedAt  time.Time `gorm:"column:updated_at"`
}
