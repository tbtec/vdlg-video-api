package dto

type OrderProduct struct {
	ID        string `json:"id"`
	ProductID string `json:"productId"`
	Quantity  int64  `json:"quantity"`
}
