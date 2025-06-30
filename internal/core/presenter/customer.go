package presenter

import (
	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type CustomerPresenter struct {
}

func NewCustomerPresenter() *CustomerPresenter {
	return &CustomerPresenter{}
}

func (presenter *CustomerPresenter) BuildCustomerCreateResponse(customer entity.Customer) dto.Customer {
	return dto.Customer{
		CustomerId:     customer.ID,
		Name:           customer.Name,
		DocumentNumber: customer.DocumentNumber,
		Email:          customer.Email,
		CreatedAt:      customer.CreatedAt,
		UpdatedAt:      customer.UpdatedAt,
	}
}

func (presenter *CustomerPresenter) BuildCustomerContentResponse(customer entity.Customer) dto.CustomerContent {
	response := dto.Customer{}

	response = presenter.BuildCustomerCreateResponse(customer)

	return dto.CustomerContent{Content: response}
}
