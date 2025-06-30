package external

type PaymentConfig struct {
	Url   string
	Token string
}

type PaymentRequest struct {
	ExternalReference string  `json:"external_reference"`
	Title             string  `json:"title"`
	Description       string  `json:"description"`
	NotificationURL   string  `json:"notification_url"`
	TotalAmount       float64 `json:"total_amount"`
	Item              []Item  `json:"items"`
	Sponsor           Sponsor `json:"sponsor"`
	CashOut           CashOut `json:"cash_out"`
}
type Sponsor struct {
	ID int `json:"id"`
}

type CashOut struct {
	Amount int `json:"amount"`
}
type Item struct {
	SkuNumber   string  `json:"sku_number"`
	Category    string  `json:"category"`
	Title       string  `json:"title"`
	Description string  `json:"description"`
	UnitPrice   float64 `json:"unit_price"`
	Quantity    int64   `json:"quantity"`
	UnitMeasure string  `json:"unit_measure"`
	TotalAmount float64 `json:"total_amount"`
}

type PaymentResponse struct {
	InStoreOrderId string `json:"in_store_order_id"`
	QRData         string `json:"qr_data"`
}
