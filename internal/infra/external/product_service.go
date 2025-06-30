package external

import (
	"context"
	"log/slog"

	"github.com/go-resty/resty/v2"
	"github.com/tbtec/tremligeiro/internal/infra/httpclient"
)

type IProductService interface {
	FindOne(ctx context.Context, id string) (*ProductResponse, error)
}

type ProductService struct {
	httpclient *resty.Client
	config     ProductConfig
}

func NewProductService(config ProductConfig) IProductService {
	return &ProductService{
		config:     config,
		httpclient: httpclient.New(),
	}
}

func (service *ProductService) FindOne(ctx context.Context, id string) (*ProductResponse, error) {
	productResponse := ProductResponse{}

	url := service.config.Url
	path := "/api/v1/product" + "/" + id

	slog.InfoContext(ctx, "ProductService - FindOne", "url", url, "path", path)

	response, err := service.httpclient.R().
		SetHeader("Content-Type", "application/json").
		SetResult(&productResponse).
		Get(url + path)
	if err != nil {
		return &productResponse, err
	}
	slog.Info("Response", "status_code", response.StatusCode(), "body", response.Body())

	if response.StatusCode() != 200 {
		return &ProductResponse{}, response.Error().(error)
	}

	return &productResponse, nil
}
