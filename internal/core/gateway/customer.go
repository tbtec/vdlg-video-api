package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
	"github.com/tbtec/tremligeiro/internal/infra/external"
)

type CustomerGateway struct {
	customerRepository repository.ICustomerRepository
	customerService    external.ICustomerService
}

func NewCustomerGateway(customerRepository repository.ICustomerRepository,
	customerService external.ICustomerService) *CustomerGateway {
	return &CustomerGateway{
		customerRepository: customerRepository,
		customerService:    customerService,
	}
}

func (gtw *CustomerGateway) FindByDocumentNumber(ctx context.Context, documentNumber string) (*entity.Customer, error) {
	customerModel, err := gtw.customerRepository.FindByDocumentNumber(ctx, documentNumber)
	if err != nil {
		return nil, err
	}

	customer := entity.Customer{
		ID:             customerModel.ID,
		Name:           customerModel.Name,
		DocumentNumber: customerModel.DocumentNumber,
		Email:          customerModel.Email,
		CreatedAt:      customerModel.CreatedAt,
		UpdatedAt:      customerModel.UpdatedAt,
	}

	return &customer, nil
}

func (gtw *CustomerGateway) FindOne(ctx context.Context, id string) (*entity.Customer, error) {
	customer := entity.Customer{}

	customerResponse, err := gtw.customerService.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	if customerResponse != nil {
		customer = entity.Customer{
			ID:             customerResponse.Content.CustomerId,
			Name:           customerResponse.Content.Name,
			DocumentNumber: customerResponse.Content.DocumentNumber,
			Email:          customerResponse.Content.Email,
			CreatedAt:      customerResponse.Content.CreatedAt,
			UpdatedAt:      customerResponse.Content.UpdatedAt,
		}
		return &customer, nil
	}

	return nil, nil
}
