package external

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPaymentConfig_Fields(t *testing.T) {
	cfg := PaymentConfig{Url: "http://pay.api", Token: "token123"}
	assert.Equal(t, "http://pay.api", cfg.Url)
	assert.Equal(t, "token123", cfg.Token)
}

func TestPaymentRequest_JSONMarshalling(t *testing.T) {
	req := PaymentRequest{
		ExternalReference: "order-1",
		Title:             "Pedido",
		Description:       "Pagamento do pedido",
		NotificationURL:   "http://notify",
		TotalAmount:       150.0,
		Item: []Item{
			{
				SkuNumber:   "sku-1",
				Category:    "cat",
				Title:       "Produto",
				Description: "desc",
				UnitPrice:   50.0,
				Quantity:    3,
				UnitMeasure: "un",
				TotalAmount: 150.0,
			},
		},
		Sponsor: Sponsor{ID: 42},
		CashOut: CashOut{Amount: 10},
	}
	data, err := json.Marshal(req)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"external_reference":"order-1"`)
	assert.Contains(t, string(data), `"items"`)
	assert.Contains(t, string(data), `"id":42`)
	assert.Contains(t, string(data), `"amount":10`)
}

func TestSponsor_Fields(t *testing.T) {
	s := Sponsor{ID: 7}
	assert.Equal(t, 7, s.ID)
}

func TestCashOut_Fields(t *testing.T) {
	c := CashOut{Amount: 99}
	assert.Equal(t, 99, c.Amount)
}

func TestItem_JSONMarshalling(t *testing.T) {
	item := Item{
		SkuNumber:   "sku-2",
		Category:    "cat2",
		Title:       "Produto2",
		Description: "desc2",
		UnitPrice:   25.0,
		Quantity:    2,
		UnitMeasure: "kg",
		TotalAmount: 50.0,
	}
	data, err := json.Marshal(item)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"sku_number":"sku-2"`)
	assert.Contains(t, string(data), `"unit_measure":"kg"`)
}

func TestPaymentResponse_JSONMarshalling(t *testing.T) {
	resp := PaymentResponse{
		InStoreOrderId: "store-123",
		QRData:         "qr-code-data",
	}
	data, err := json.Marshal(resp)
	assert.NoError(t, err)
	assert.Contains(t, string(data), `"in_store_order_id":"store-123"`)
	assert.Contains(t, string(data), `"qr_data":"qr-code-data"`)
}
