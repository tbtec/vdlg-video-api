package external

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCustomerConfig_Fields(t *testing.T) {
	cfg := CustomerConfig{Url: "http://customer.api"}
	assert.Equal(t, "http://customer.api", cfg.Url)
}

func TestCustomerResponse_JSONMarshalling(t *testing.T) {
	now := time.Now()
	resp := CustomerResponse{
		CustomerId:     "c1",
		Name:           "João",
		DocumentNumber: "12345678900",
		Email:          "joao@email.com",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	data, err := json.Marshal(resp)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":"c1"`)
	assert.Contains(t, string(data), `"name":"João"`)
	assert.Contains(t, string(data), `"documentNumber":"12345678900"`)
	assert.Contains(t, string(data), `"email":"joao@email.com"`)

	var unmarshalled CustomerResponse
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
}

func TestCustomerContent_JSONMarshalling(t *testing.T) {
	now := time.Now()
	resp := CustomerResponse{
		CustomerId:     "c2",
		Name:           "Maria",
		DocumentNumber: "98765432100",
		Email:          "maria@email.com",
		CreatedAt:      now,
		UpdatedAt:      now,
	}
	content := CustomerContent{Content: resp}
	data, err := json.Marshal(content)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"id":"c2"`)
	assert.Contains(t, string(data), `"name":"Maria"`)

	var unmarshalled CustomerContent
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
}
