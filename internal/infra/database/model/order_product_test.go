package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrderProduct_Fields(t *testing.T) {
	now := time.Now()
	op := OrderProduct{
		ID:          "op1",
		OrderID:     "order1",
		ProductID:   "prod1",
		Quantity:    2,
		Amount:      10.5,
		TotalAmount: 21.0,
		CreatedAt:   now,
	}
	assert.Equal(t, "op1", op.ID)
	assert.Equal(t, "order1", op.OrderID)
	assert.Equal(t, "prod1", op.ProductID)
	assert.Equal(t, int64(2), op.Quantity)
	assert.Equal(t, 10.5, op.Amount)
	assert.Equal(t, 21.0, op.TotalAmount)
	assert.Equal(t, now, op.CreatedAt)
}

func TestOrderProduct_JSONMarshalling(t *testing.T) {
	now := time.Now()
	op := OrderProduct{
		ID:          "op2",
		OrderID:     "order2",
		ProductID:   "prod2",
		Quantity:    5,
		Amount:      7.0,
		TotalAmount: 35.0,
		CreatedAt:   now,
	}
	data, err := json.Marshal(op)
	assert.NoError(t, err)

	var unmarshalled OrderProduct
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
}
