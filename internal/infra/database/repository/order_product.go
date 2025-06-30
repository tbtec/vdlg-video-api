package repository

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/infra/database"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

type IOrderProductRepository interface {
	Create(ctx context.Context, product *model.OrderProduct) error
	FindByOrderId(ctx context.Context, orderId string) ([]model.OrderProduct, error)
}

type OrderProductRepository struct {
	database database.RDBMS
}

func NewOrderProductRepository(database database.RDBMS) IOrderProductRepository {
	return &OrderProductRepository{
		database: database,
	}
}

func (repository *OrderProductRepository) Create(ctx context.Context, product *model.OrderProduct) error {
	result := repository.database.DB.WithContext(ctx).Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *OrderProductRepository) FindByOrderId(ctx context.Context, orderId string) ([]model.OrderProduct, error) {
	orderProducts := []model.OrderProduct{}

	result := repository.database.DB.WithContext(ctx).Where("order_id = ?", orderId).Find(&orderProducts)

	if result.Error != nil {
		return nil, result.Error
	}

	return orderProducts, nil
}
