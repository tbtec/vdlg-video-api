package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/dto"
)

func TestNewOrder_WithCustomer(t *testing.T) {
	input := dto.CreateOrder{CustomerId: "c1"}
	order := NewOrder(input, true)
	assert.NotEmpty(t, order.ID)
	assert.NotNil(t, order.CustomerId)
	assert.Equal(t, "c1", *order.CustomerId)
	assert.Equal(t, OrderStatusPending, order.Status)
	assert.WithinDuration(t, time.Now().UTC(), order.CreatedAt, time.Second)
	assert.WithinDuration(t, time.Now().UTC(), order.UpdatedAt, time.Second)
}

func TestNewOrder_WithoutCustomer(t *testing.T) {
	input := dto.CreateOrder{CustomerId: ""}
	order := NewOrder(input, false)
	assert.NotEmpty(t, order.ID)
	assert.Nil(t, order.CustomerId)
	assert.Equal(t, OrderStatusPending, order.Status)
}

func TestOrder_SetTotalAmount(t *testing.T) {
	order := Order{}
	order.SetTotalAmount(123.45)
	assert.Equal(t, 123.45, order.TotalAmount)
}

func TestOrder_SetStatus(t *testing.T) {
	order := Order{}
	order.SetStatus(OrderStatusReady)
	assert.Equal(t, OrderStatusReady, order.Status)
}

func TestOrder_ValidateStatus(t *testing.T) {
	order := Order{}

	assert.True(t, order.ValidateStatus(OrderStatusPending, OrderStatusReceived))
	assert.True(t, order.ValidateStatus(OrderStatusPending, OrderStatusExpired))
	assert.True(t, order.ValidateStatus(OrderStatusPending, OrderStatusInPreparation))
	assert.False(t, order.ValidateStatus(OrderStatusPending, OrderStatusFinalized))

	assert.True(t, order.ValidateStatus(OrderStatusReceived, OrderStatusInPreparation))
	assert.False(t, order.ValidateStatus(OrderStatusReceived, OrderStatusReady))

	assert.True(t, order.ValidateStatus(OrderStatusInPreparation, OrderStatusReady))
	assert.False(t, order.ValidateStatus(OrderStatusInPreparation, OrderStatusFinalized))

	assert.True(t, order.ValidateStatus(OrderStatusReady, OrderStatusFinalized))
	assert.False(t, order.ValidateStatus(OrderStatusReady, OrderStatusPending))
}
