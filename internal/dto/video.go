package dto

import "time"

type CreateVideo struct {
	CustomerId string `json:"-" validate:"required"`
	FileName   string `json:"fileName" validate:"required"`
}

type UpdateVideo struct {
	InputMessage  *InputMessage
	OutputMessage *OutputMessage
}
type Video struct {
	ID             string    `json:"id"`
	CustomerId     string    `json:"customerId"`
	Status         string    `json:"status"`
	FileNameInput  string    `json:"fileNameInput"`
	FileNameOutput string    `json:"fileNameOutput"`
	UploadUrl      *string   `json:"uploadUrl,omitempty"`
	DownloadUrl    *string   `json:"downloadUrl,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

type VideoContent struct {
	Content []Video `json:"content"`
}
