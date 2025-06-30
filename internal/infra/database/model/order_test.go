package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestOrder_Fields(t *testing.T) {
	now := time.Now()
	customerId := "c1"
	order := Order{
		ID:          "o1",
		CustomerId:  &customerId,
		Status:      "PENDING",
		TotalAmount: 100.50,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	assert.Equal(t, "o1", order.ID)
	assert.NotNil(t, order.CustomerId)
	assert.Equal(t, "c1", *order.CustomerId)
	assert.Equal(t, "PENDING", order.Status)
	assert.Equal(t, 100.50, order.TotalAmount)
	assert.Equal(t, now, order.CreatedAt)
	assert.Equal(t, now, order.UpdatedAt)
}

func TestOrder_JSONMarshalling(t *testing.T) {
	now := time.Now()
	customerId := "c2"
	order := Order{
		ID:          "o2",
		CustomerId:  &customerId,
		Status:      "FINALIZED",
		TotalAmount: 200.00,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	data, err := json.Marshal(order)
	assert.NoError(t, err)

	var unmarshalled Order
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
	assert.Equal(t, order.ID, unmarshalled.ID)
	assert.NotNil(t, unmarshalled.CustomerId)
	assert.Equal(t, *order.CustomerId, *unmarshalled.CustomerId)
	assert.Equal(t, order.Status, unmarshalled.Status)
	assert.Equal(t, order.TotalAmount, unmarshalled.TotalAmount)
}
