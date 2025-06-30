package repository

import (
	"context"
	"fmt"

	"github.com/tbtec/tremligeiro/internal/infra/database"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

type IProductRepository interface {
	Create(ctx context.Context, product *model.Product) error
	FindOne(ctx context.Context, id string) (*model.Product, error)
	FindByCategory(ctx context.Context, id int) (*[]model.Product, error)
	DeleteById(ctx context.Context, id string) (*model.Product, error)
	UpdateById(ctx context.Context, product *model.Product) error
}

type ProductRepository struct {
	database database.RDBMS
}

func NewProductRepository(database database.RDBMS) IProductRepository {
	return &ProductRepository{
		database: database,
	}
}

func (repository *ProductRepository) Create(ctx context.Context, product *model.Product) error {

	result := repository.database.DB.WithContext(ctx).Create(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *ProductRepository) FindOne(ctx context.Context, id string) (*model.Product, error) {
	product := &model.Product{}

	result := repository.database.DB.WithContext(ctx).Where("product_id = ?", id).First(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repository *ProductRepository) FindByCategory(ctx context.Context, id int) (*[]model.Product, error) {
	product := &[]model.Product{}

	result := repository.database.DB.WithContext(ctx).Where("category_id = ?", id).Find(&product)

	if result.Error != nil {
		return nil, result.Error
	}

	return product, nil
}

func (repository *ProductRepository) DeleteById(ctx context.Context, id string) (*model.Product, error) {
	product := &model.Product{
		ID: id,
	}

	result := repository.database.DB.WithContext(ctx).Delete(&product)

	if result.Error != nil {
		return nil, result.Error
	} else if result.RowsAffected < 1 {
		return product, fmt.Errorf("record not found")
	}

	return product, nil
}

func (repository *ProductRepository) UpdateById(ctx context.Context, product *model.Product) error {

	result := repository.database.DB.WithContext(ctx).Save(&product)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
