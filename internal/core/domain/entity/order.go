package entity

import (
	"time"

	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/types/ulid"
)

type Order struct {
	ID          string
	CustomerId  *string
	Status      OrderStatus
	TotalAmount float64
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

type OrderStatus string

const (
	OrderStatusPending       OrderStatus = "PENDING"
	OrderStatusReceived      OrderStatus = "RECEIVED"
	OrderStatusInPreparation OrderStatus = "IN_PREPARATION"
	OrderStatusReady         OrderStatus = "READY"
	OrderStatusFinalized     OrderStatus = "FINALIZED"
	OrderStatusExpired       OrderStatus = "EXPIRED"
)

func NewOrder(createOrder dto.CreateOrder, customerPresent bool) Order {
	order := Order{
		ID:        ulid.NewUlid().String(),
		Status:    OrderStatusPending,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	if customerPresent {
		order.CustomerId = &createOrder.CustomerId
	}

	return order
}

func (order *Order) SetTotalAmount(amount float64) {
	order.TotalAmount = amount
}

func (order *Order) SetStatus(status OrderStatus) {
	order.Status = status
}

func (order *Order) ValidateStatus(currentStatus OrderStatus, newStatus OrderStatus) bool {
	switch currentStatus {
	case OrderStatusPending:
		return newStatus == OrderStatusReceived || newStatus == OrderStatusExpired || newStatus == OrderStatusInPreparation
	case OrderStatusReceived:
		return newStatus == OrderStatusInPreparation
	case OrderStatusInPreparation:
		return newStatus == OrderStatusReady
	case OrderStatusReady:
		return newStatus == OrderStatusFinalized
	}

	return false
}
