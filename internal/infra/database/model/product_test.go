package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProduct_Fields(t *testing.T) {
	now := time.Now()
	p := Product{
		ID:          "prod1",
		Name:        "Coca-Cola",
		Description: "Refrigerante",
		CategoryId:  1,
		Amount:      5.50,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	assert.Equal(t, "prod1", p.ID)
	assert.Equal(t, "Coca-Cola", p.Name)
	assert.Equal(t, "Refrigerante", p.Description)
	assert.Equal(t, 1, p.CategoryId)
	assert.Equal(t, 5.50, p.Amount)
	assert.Equal(t, now, p.CreatedAt)
	assert.Equal(t, now, p.UpdatedAt)
}

func TestProduct_JSONMarshalling(t *testing.T) {
	now := time.Now()
	p := Product{
		ID:          "prod2",
		Name:        "Pepsi",
		Description: "Refrigerante de cola",
		CategoryId:  2,
		Amount:      4.99,
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	data, err := json.Marshal(p)
	assert.NoError(t, err)

	var unmarshalled Product
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
}
