package model

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestPayment_Fields(t *testing.T) {
	now := time.Now()
	p := Payment{
		ID:         "pay1",
		OrderId:    "order1",
		Status:     "PENDING",
		QrData:     "qrdata",
		ExternalId: "ext1",
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	assert.Equal(t, "pay1", p.ID)
	assert.Equal(t, "order1", p.OrderId)
	assert.Equal(t, "PENDING", p.Status)
	assert.Equal(t, "qrdata", p.QrData)
	assert.Equal(t, "ext1", p.ExternalId)
	assert.Equal(t, now, p.CreatedAt)
	assert.Equal(t, now, p.UpdatedAt)
}

func TestPayment_JSONMarshalling(t *testing.T) {
	now := time.Now()
	p := Payment{
		ID:         "pay2",
		OrderId:    "order2",
		Status:     "FINALIZED",
		QrData:     "qrdata2",
		ExternalId: "ext2",
		CreatedAt:  now,
		UpdatedAt:  now,
	}
	data, err := json.Marshal(p)
	assert.NoError(t, err)

	var unmarshalled Payment
	err = json.Unmarshal(data, &unmarshalled)
	assert.NoError(t, err)
}
