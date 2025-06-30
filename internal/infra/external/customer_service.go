package external

import (
	"context"

	"github.com/go-resty/resty/v2"
	"github.com/tbtec/tremligeiro/internal/infra/httpclient"
)

type ICustomerService interface {
	FindOne(ctx context.Context, id string) (*CustomerContent, error)
}

type CustomerService struct {
	httpclient *resty.Client
	config     CustomerConfig
}

func NewCustomerService(config CustomerConfig) ICustomerService {
	return &CustomerService{
		config:     config,
		httpclient: httpclient.New(),
	}
}

func (service *CustomerService) FindOne(ctx context.Context, id string) (*CustomerContent, error) {
	customerResponse := CustomerContent{}

	url := service.config.Url
	path := "/api/v1/customer" + "/" + id

	response, err := service.httpclient.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&customerResponse).
		Get(url + path)
	if err != nil {
		return &customerResponse, err
	}

	if response.StatusCode() != 200 {
		return &CustomerContent{}, response.Error().(error)
	}

	return &customerResponse, nil
}
