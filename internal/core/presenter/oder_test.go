package presenter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
)

func TestBuildOrderCreateResponse(t *testing.T) {
	presenter := NewOrderPresenter()
	now := time.Now()
	customerId := "c1"
	order := entity.Order{
		ID:          "o1",
		CustomerId:  &customerId,
		Status:      "PENDING",
		TotalAmount: 100.0,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	paymentId := "pay-123"
	resp := presenter.BuildOrderCreateResponse(order, &paymentId)
	assert.Equal(t, order.ID, resp.ID)
	assert.Equal(t, order.CustomerId, resp.CustomerId)
	assert.Equal(t, string(order.Status), resp.Status)
	assert.Equal(t, order.TotalAmount, resp.TotalAmount)
	assert.Equal(t, order.CreatedAt, resp.CreatedAt)
	assert.Equal(t, order.UpdatedAt, resp.UpdatedAt)
	assert.NotNil(t, resp.MetaData.PaymentId)
	assert.Equal(t, paymentId, *resp.MetaData.PaymentId)
}

func TestBuildOrderContentResponse(t *testing.T) {
	presenter := NewOrderPresenter()
	now := time.Now()
	customerId1 := "c1"
	customerId2 := "c2"
	customerId3 := "c3"
	orders := []entity.Order{
		{
			ID:          "o1",
			CustomerId:  &customerId1,
			Status:      "PENDING",
			TotalAmount: 100.0,
			CreatedAt:   now.Add(-2 * time.Hour),
			UpdatedAt:   now.Add(-2 * time.Hour),
		},
		{
			ID:          "o2",
			CustomerId:  &customerId2,
			Status:      "FINALIZED",
			TotalAmount: 200.0,
			CreatedAt:   now.Add(-1 * time.Hour),
			UpdatedAt:   now.Add(-1 * time.Hour),
		},
		{
			ID:          "o3",
			CustomerId:  &customerId3,
			Status:      "PENDING",
			TotalAmount: 300.0,
			CreatedAt:   now,
			UpdatedAt:   now,
		},
	}
	content := presenter.BuildOrderContentResponse(orders)
	assert.Len(t, content.Content, 2)
	assert.Equal(t, "o1", content.Content[0].ID)
	assert.Equal(t, "o3", content.Content[1].ID)
}

func TestBuildOrderDetailsCreateResponse(t *testing.T) {
	presenter := NewOrderPresenter()
	now := time.Now()
	customerId := "c1"
	order := entity.Order{
		ID:          "o1",
		CustomerId:  &customerId,
		Status:      "PENDING",
		TotalAmount: 100.0,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	orderProducts := []entity.OrderProduct{
		{ID: "op1", ProductID: "p1", Quantity: 2},
		{ID: "op2", ProductID: "p2", Quantity: 1},
	}
	details := presenter.BuildOrderDetailsCreateResponse(order, orderProducts)
	assert.Equal(t, order.ID, details.ID)
	assert.Equal(t, order.CustomerId, details.CustomerId)
	assert.Equal(t, string(order.Status), details.Status)
	assert.Equal(t, order.TotalAmount, details.TotalAmount)
	assert.Equal(t, order.CreatedAt, details.CreatedAt)
	assert.Equal(t, order.UpdatedAt, details.UpdatedAt)
	assert.Len(t, details.OrderProducts, 2)
	assert.Equal(t, "op1", details.OrderProducts[0].ID)
	assert.Equal(t, "p1", details.OrderProducts[0].ProductID)
	assert.Equal(t, int64(2), details.OrderProducts[0].Quantity)
	assert.Equal(t, "op2", details.OrderProducts[1].ID)
	assert.Equal(t, "p2", details.OrderProducts[1].ProductID)
	assert.Equal(t, int64(1), details.OrderProducts[1].Quantity)
}
