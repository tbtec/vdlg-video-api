package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
	"github.com/tbtec/tremligeiro/internal/infra/external"
)

type ProductGateway struct {
	productRepository repository.IProductRepository
	productService    external.IProductService
}

func NewProductGateway(productRepository repository.IProductRepository,
	productService external.IProductService) *ProductGateway {
	return &ProductGateway{
		productRepository: productRepository,
		productService:    productService,
	}
}

func (gtw *ProductGateway) FindOne(ctx context.Context, id string) (*entity.Product, error) {

	productResponse, err := gtw.productService.FindOne(ctx, id)
	if productResponse == nil {
		return nil, err
	}

	product := entity.Product{
		ID:          productResponse.ProductId,
		Name:        productResponse.Name,
		Description: productResponse.Description,
		Amount:      productResponse.Amount,
		CategoryId:  productResponse.Category.ID,
		CreatedAt:   productResponse.CreatedAt,
		UpdatedAt:   productResponse.UpdatedAt,
	}

	return &product, nil
}
