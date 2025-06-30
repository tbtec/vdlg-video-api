package repository

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/infra/database"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

type ICustomerRepository interface {
	Create(ctx context.Context, customer *model.Customer) error
	FindOne(ctx context.Context, id string) (*model.Customer, error)
	FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Customer, error)
}

type CustomerRepository struct {
	database database.RDBMS
}

func NewCustomerRepository(database database.RDBMS) ICustomerRepository {
	return &CustomerRepository{
		database: database,
	}
}

func (repository *CustomerRepository) Create(ctx context.Context, customer *model.Customer) error {

	result := repository.database.DB.WithContext(ctx).Create(&customer)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *CustomerRepository) FindOne(ctx context.Context, id string) (*model.Customer, error) {
	customer := &model.Customer{}

	result := repository.database.DB.WithContext(ctx).Where("customer_id = ?", id).First(&customer)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, nil
		}
		return nil, result.Error
	}

	return customer, nil
}

func (repository *CustomerRepository) FindByDocumentNumber(ctx context.Context, documentNumber string) (*model.Customer, error) {

	customer := &model.Customer{}

	result := repository.database.DB.Where("document_number = ?", documentNumber).First(&customer)

	if result.Error != nil {
		return nil, result.Error
	}

	return customer, nil
}
