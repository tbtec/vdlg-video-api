package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProduct_Fields(t *testing.T) {
	now := time.Now()
	p := Product{
		ID:          "p1",
		Name:        "Coca-Cola",
		Description: "Refrigerante",
		CategoryId:  1,
		Amount:      5.50,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	assert.Equal(t, "p1", p.ID)
	assert.Equal(t, "Coca-Cola", p.Name)
	assert.Equal(t, "Refrigerante", p.Description)
	assert.Equal(t, 1, p.CategoryId)
	assert.Equal(t, 5.50, p.Amount)
	assert.Equal(t, now, p.CreatedAt)
	assert.Equal(t, now, p.UpdatedAt)
}
