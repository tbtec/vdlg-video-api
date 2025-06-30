package dto

type CreateVideo struct {
	CustomerId string `json:"-" validate:"required"`
	FileName   string `json:"fileName" validate:"required"`
}
