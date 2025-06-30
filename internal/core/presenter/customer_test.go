package presenter

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
)

func TestBuildCustomerCreateResponse(t *testing.T) {
	presenter := NewCustomerPresenter()
	now := time.Now()
	customer := entity.Customer{
		ID:             "c1",
		Name:           "John Doe",
		DocumentNumber: "123456789",
		Email:          "john@example.com",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	resp := presenter.BuildCustomerCreateResponse(customer)
	assert.Equal(t, customer.ID, resp.CustomerId)
	assert.Equal(t, customer.Name, resp.Name)
	assert.Equal(t, customer.DocumentNumber, resp.DocumentNumber)
	assert.Equal(t, customer.Email, resp.Email)
	assert.Equal(t, customer.CreatedAt, resp.CreatedAt)
	assert.Equal(t, customer.UpdatedAt, resp.UpdatedAt)
}

func TestBuildCustomerContentResponse(t *testing.T) {
	presenter := NewCustomerPresenter()
	now := time.Now()
	customer := entity.Customer{
		ID:             "c2",
		Name:           "Jane Doe",
		DocumentNumber: "987654321",
		Email:          "jane@example.com",
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	content := presenter.BuildCustomerContentResponse(customer)
	assert.Equal(t, customer.ID, content.Content.CustomerId)
	assert.Equal(t, customer.Name, content.Content.Name)
	assert.Equal(t, customer.DocumentNumber, content.Content.DocumentNumber)
	assert.Equal(t, customer.Email, content.Content.Email)
	assert.Equal(t, customer.CreatedAt, content.Content.CreatedAt)
	assert.Equal(t, customer.UpdatedAt, content.Content.UpdatedAt)
}
