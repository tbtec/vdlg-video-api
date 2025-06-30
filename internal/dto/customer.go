package dto

import "time"

type CreateCustomer struct {
	Name           string `json:"name" validate:"required"`
	DocumentNumber string `json:"documentNumber" validate:"required"`
	Email          string `json:"email" validate:"required"`
}

type UpdateCustomer struct {
	CustomerId     string
	Name           string
	DocumentNumber string
	Email          string
	UpdatedAt      time.Time
}

type FindCustomer struct {
	DocumentNumber string `json:"documentNumber" validate:"required"`
}

type Customer struct {
	CustomerId     string    `json:"id"`
	Name           string    `json:"name"`
	DocumentNumber string    `json:"documentNumber"`
	Email          string    `json:"email"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type CustomerContent struct {
	Content Customer `json:"content"`
}
