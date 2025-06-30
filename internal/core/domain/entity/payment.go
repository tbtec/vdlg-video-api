package entity

import (
	"time"

	"github.com/tbtec/tremligeiro/internal/types/ulid"
)

type Payment struct {
	ID         string
	OrderId    string
	Status     PaymentStatus
	QrData     string
	ExternalId string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type PaymentStatus string

const (
	PaymentStatusPending       PaymentStatus = "PENDING"
	PaymentStatusAuthorized    PaymentStatus = "AUTHORIZED"
	PaymentStatusNotAuthorized PaymentStatus = "NOT_AUTHORIZED"
)

func NewPayment(orderId string) Payment {
	return Payment{
		ID:        ulid.NewUlid().String(),
		OrderId:   orderId,
		Status:    PaymentStatusPending,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}
}

func (payment *Payment) SetStatus(status PaymentStatus) {
	payment.Status = status
}

func (payment *Payment) IsFinished() bool {
	return payment.Status == PaymentStatusAuthorized || payment.Status == PaymentStatusNotAuthorized
}
