package model

import "time"

type Customer struct {
	ID             string    `gorm:"column:customer_id;primaryKey"`
	Name           string    `gorm:"column:name"`
	DocumentNumber string    `gorm:"column:document_number"`
	Email          string    `gorm:"column:email"`
	CreatedAt      time.Time `gorm:"column:created_at"`
	UpdatedAt      time.Time `gorm:"column:updated_at"`
}
