package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type CreateOrderController struct {
	usc *usecase.UscCreateOrder
}

func NewCreateOrderController(container *container.Container) *CreateOrderController {
	return &CreateOrderController{
		usc: usecase.NewUseCaseCreateOrder(
			gateway.NewOrderGateway(container.OrderRepository),
			gateway.NewCustomerGateway(container.CustomerRepository, container.CustomerService),
		),
	}
}

func (ctl *CreateOrderController) Execute(ctx context.Context, input dto.CreateOrder) (dto.Order, error) {
	return ctl.usc.Create(ctx, input)
}
