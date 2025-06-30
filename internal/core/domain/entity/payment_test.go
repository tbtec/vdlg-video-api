package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewPayment(t *testing.T) {
	orderId := "order-123"
	p := NewPayment(orderId)
	assert.NotEmpty(t, p.ID)
	assert.Equal(t, orderId, p.OrderId)
	assert.Equal(t, PaymentStatusPending, p.Status)
	assert.WithinDuration(t, time.Now().UTC(), p.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now().UTC(), p.UpdatedAt, time.Second)
}

func TestPayment_SetStatus(t *testing.T) {
	p := NewPayment("order-456")
	p.SetStatus(PaymentStatusAuthorized)
	assert.Equal(t, PaymentStatusAuthorized, p.Status)
	p.SetStatus(PaymentStatusNotAuthorized)
	assert.Equal(t, PaymentStatusNotAuthorized, p.Status)
}

func TestPayment_IsFinished(t *testing.T) {
	p := NewPayment("order-789")
	assert.False(t, p.IsFinished())

	p.SetStatus(PaymentStatusAuthorized)
	assert.True(t, p.IsFinished())

	p.SetStatus(PaymentStatusNotAuthorized)
	assert.True(t, p.IsFinished())

	p.SetStatus(PaymentStatusPending)
	assert.False(t, p.IsFinished())
}
