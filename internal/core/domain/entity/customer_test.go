package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomer_Fields(t *testing.T) {
	now := time.Now()
	c := Customer{
		ID:             "c1",
		Name:           "João",
		DocumentNumber: "12345678900",
		Email:          "joao@email.com",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	assert.Equal(t, "c1", c.ID)
	assert.Equal(t, "João", c.Name)
	assert.Equal(t, "12345678900", c.DocumentNumber)
	assert.Equal(t, "joao@email.com", c.Email)
	assert.Equal(t, now, c.CreatedAt)
	assert.Equal(t, now, c.UpdatedAt)
}
