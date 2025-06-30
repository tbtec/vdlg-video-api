package dto

type MetaDataContent struct {
	Content MetaData `json:"metadata"`
}

type MetaData struct {
	PaymentId         *string `json:"paymentId"`
	PaymentWebHookUrl *string `json:"paymentWebhookUrl"`
}
