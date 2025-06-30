package model

import "time"

type Order struct {
	ID          string    `gorm:"column:order_id;primaryKey"`
	CustomerId  *string   `gorm:"column:customer_id"`
	Status      string    `gorm:"column:status"`
	TotalAmount float64   `gorm:"column:total_amount"`
	CreatedAt   time.Time `gorm:"column:created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at"`
}
