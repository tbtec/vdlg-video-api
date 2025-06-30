package entity

import (
	"time"
)

type Product struct {
	ID          string
	Name        string
	Description string
	CategoryId  int
	Amount      float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
