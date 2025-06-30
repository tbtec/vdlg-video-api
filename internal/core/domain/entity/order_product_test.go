package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestNewOrderProduct(t *testing.T) {
	orderId := "order-1"
	productId := "prod-1"
	quantity := 3
	amount := 10.5

	op := NewOrderProduct(orderId, productId, quantity, amount)

	assert.NotEmpty(t, op.ID)
	assert.Equal(t, orderId, op.OrderID)
	assert.Equal(t, productId, op.ProductID)
	assert.Equal(t, int64(quantity), op.Quantity)
	assert.Equal(t, amount, op.Amount)
	assert.Equal(t, float64(quantity)*amount, op.TotalAmount)
	assert.WithinDuration(t, time.Now().UTC(), op.CreatedAt, time.Second)
}
