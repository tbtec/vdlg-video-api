package external

import "time"

type ProductConfig struct {
	Url string
}

type ProductResponse struct {
	ProductId   string           `json:"id"`
	Name        string           `json:"name"`
	Description string           `json:"description"`
	Amount      float64          `json:"amount"`
	Category    CategoryResponse `json:"category"`
	CreatedAt   time.Time        `json:"createdAt"`
	UpdatedAt   time.Time        `json:"updatedAt"`
}

type CategoryResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}
