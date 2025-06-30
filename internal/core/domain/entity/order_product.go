package entity

import (
	"time"

	"github.com/tbtec/tremligeiro/internal/types/ulid"
)

type OrderProduct struct {
	ID          string
	OrderID     string
	ProductID   string
	Quantity    int64
	Amount      float64
	TotalAmount float64
	CreatedAt   time.Time
}

func NewOrderProduct(orderId string, productId string, quantity int, amount float64) OrderProduct {
	return OrderProduct{
		ID:          ulid.NewUlid().String(),
		OrderID:     orderId,
		ProductID:   productId,
		Quantity:    int64(quantity),
		Amount:      amount,
		TotalAmount: float64(quantity) * amount,
		CreatedAt:   time.Now().UTC(),
	}
}
