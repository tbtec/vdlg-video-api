package repository

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/infra/database"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
)

type IPaymentRepository interface {
	Create(ctx context.Context, payment *model.Payment) error
	FindOne(ctx context.Context, id string) (*model.Payment, error)
	FindByOrderId(ctx context.Context, id string) (*model.Payment, error)
	Update(ctx context.Context, payment *model.Payment) error
}

type PaymentRepository struct {
	database database.RDBMS
}

func NewPaymentRepository(database database.RDBMS) IPaymentRepository {
	return &PaymentRepository{
		database: database,
	}
}

func (repository *PaymentRepository) Create(ctx context.Context, payment *model.Payment) error {

	result := repository.database.DB.WithContext(ctx).Create(&payment)

	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (repository *PaymentRepository) FindOne(ctx context.Context, id string) (*model.Payment, error) {
	payment := &model.Payment{}

	result := repository.database.DB.WithContext(ctx).Where("payment_id = ?", id).First(&payment)

	if result.Error != nil {
		return nil, result.Error
	}

	return payment, nil
}

func (repository *PaymentRepository) FindByOrderId(ctx context.Context, id string) (*model.Payment, error) {
	payment := model.Payment{}

	result := repository.database.DB.WithContext(ctx).Where("order_id = ?", id).Order("created_at DESC").First(&payment)

	if result.Error != nil {
		if result.Error.Error() == "record not found" {
			return nil, nil
		}
		return nil, result.Error
	}

	return &payment, nil
}

func (repository *PaymentRepository) Update(ctx context.Context, payment *model.Payment) error {

	result := repository.database.DB.WithContext(ctx).Save(&payment)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
