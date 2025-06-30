package dto

type PaymentCheckout struct {
	OrderId     string                       `json:"orderId" validate:"required"`
	TotalAmount float64                      `json:"totalAmount"`
	Products    []PaymentItemCheckoutProduct `json:"products" validate:"required"`
	MetaData    MetaData                     `json:"metadata"`
}

type PaymentItemCheckoutProduct struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
