package model

import (
	"time"
)

type OrderProduct struct {
	ID          string    `gorm:"column:order_product_id;primaryKey"`
	OrderID     string    `gorm:"column:order_id"`
	ProductID   string    `gorm:"column:product_id"`
	Quantity    int64     `gorm:"column:quantity"`
	Amount      float64   `gorm:"column:amount"`
	TotalAmount float64   `gorm:"column:total_amount"`
	CreatedAt   time.Time `gorm:"column:created_at"`
}
