package repository

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/infra/database"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

type IOrderRepository interface {
	Create(ctx context.Context, order *model.Order) error
	Find(ctx context.Context) ([]model.Order, error)
	FindOne(ctx context.Context, id string) (*model.Order, error)
	Update(ctx context.Context, order *model.Order) error
}

type OrderRepository struct {
	database database.RDBMS
}

func NewOrderRepository(database database.RDBMS) IOrderRepository {
	return &OrderRepository{
		database: database,
	}
}

func (repository *OrderRepository) Create(ctx context.Context, order *model.Order) error {

	result := repository.database.DB.WithContext(ctx).Create(&order)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *OrderRepository) Find(ctx context.Context) ([]model.Order, error) {
	orders := []model.Order{}

	result := repository.database.DB.WithContext(ctx).Find(&orders)
	if result.Error != nil {
		return nil, result.Error
	}

	return orders, nil
}

func (repository *OrderRepository) FindOne(ctx context.Context, id string) (*model.Order, error) {
	order := &model.Order{}

	result := repository.database.DB.WithContext(ctx).Where("order_id = ?", id).First(&order)

	if result.Error != nil {
		return nil, result.Error
	}

	return order, nil
}

func (repository *OrderRepository) Update(ctx context.Context, order *model.Order) error {

	result := repository.database.DB.WithContext(ctx).Save(&order)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
