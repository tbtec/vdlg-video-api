package external

import (
	"context"
	"log/slog"

	"github.com/go-resty/resty/v2"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/httpclient"
)

type IPaymentService interface {
	RequestPayment(ctx context.Context, request dto.PaymentCheckout) (PaymentResponse, error)
}

type PaymentService struct {
	httpclient *resty.Client
	config     PaymentConfig
}

func NewPaymentService(config PaymentConfig) IPaymentService {
	return &PaymentService{
		config:     config,
		httpclient: httpclient.New(),
	}
}

func (service *PaymentService) RequestPayment(ctx context.Context, request dto.PaymentCheckout) (PaymentResponse, error) {
	paymentResponse := PaymentResponse{}

	url := service.config.Url + "/api/v1/payment"

	slog.InfoContext(ctx, "PaymentService - Request Payment", "url", url)

	response, err := service.httpclient.R().
		SetHeader("Content-Type", "application/json").
		SetBody(request).
		SetResult(&paymentResponse).
		Post(url)
	if err != nil {
		return paymentResponse, err
	}

	if response.StatusCode() != 200 {
		return PaymentResponse{}, response.Error().(error)
	}

	return paymentResponse, nil
}
