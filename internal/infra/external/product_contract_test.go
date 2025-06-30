package external

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestProductConfig_Fields(t *testing.T) {
	cfg := ProductConfig{Url: "http://product.api"}
	assert.Equal(t, "http://product.api", cfg.Url)
}

func TestProductResponse_JSONMarshalling(t *testing.T) {
	now := time.Now()
	resp := ProductResponse{
		ProductId:   "p1",
		Name:        "Coca-Cola",
		Description: "Refrigerante",
		Amount:      5.5,
		Category:    CategoryResponse{ID: 1, Name: "Bebidas"},
		CreatedAt:   now,
		UpdatedAt:   now,
	}
	data, err := json.Marshal(resp)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":"p1"`)
	assert.Contains(t, string(data), `"name":"Coca-Cola"`)
	assert.Contains(t, string(data), `"description":"Refrigerante"`)
	assert.Contains(t, string(data), `"amount":5.5`)
	assert.Contains(t, string(data), `"category":{"id":1,"name":"Bebidas"}`)

	var unmarshalled ProductResponse
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
}

func TestCategoryResponse_JSONMarshalling(t *testing.T) {
	cat := CategoryResponse{ID: 2, Name: "Alimentos"}
	data, err := json.Marshal(cat)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":2`)
	assert.Contains(t, string(data), `"name":"Alimentos"`)

	var unmarshalled CategoryResponse
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
}
