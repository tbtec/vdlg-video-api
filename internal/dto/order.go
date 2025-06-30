package dto

import "time"

type CreateOrder struct {
	CustomerId string `json:"customerId"`
}

type Order struct {
	ID          string    `json:"id"`
	CustomerId  *string   `json:"customerId"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"totalAmount"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
	MetaData    MetaData  `json:"metadata"`
}

type OrderCheckout struct {
	OrderId  string                 `json:"orderId" validate:"required"`
	Products []OrderCheckoutProduct `json:"products" validate:"required"`
	MetaData MetaData               `json:"metadata"`
}

type OrderCheckoutProduct struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}
type OrderContent struct {
	Content []Order `json:"content"`
}

type OrderDetails struct {
	ID            string         `json:"id"`
	CustomerId    *string        `json:"customerId"`
	Status        string         `json:"status"`
	TotalAmount   float64        `json:"totalAmount"`
	CreatedAt     time.Time      `json:"createdAt"`
	UpdatedAt     time.Time      `json:"updatedAt"`
	OrderProducts []OrderProduct `json:"orderProducts"`
}

type UpdateOrder struct {
	Status string `json:"status" validate:"oneof='IN_PREPARATION' 'READY' 'FINALIZED'"`
}

type OrderEvent struct {
	ID     string `json:"id"`
	Status string `json:"status"`
}
