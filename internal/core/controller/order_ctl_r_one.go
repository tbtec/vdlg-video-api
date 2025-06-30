package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type FindOneOrderController struct {
	usc *usecase.UscFindOneOrder
}

func NewFindOneOrderController(container *container.Container) *FindOneOrderController {
	return &FindOneOrderController{
		usc: usecase.NewUscFindOneOrder(
			gateway.NewOrderGateway(container.OrderRepository),
			gateway.NewOrderProductGateway(container.OrderProductRepository),
			gateway.NewPaymentGateway(container.PaymentService, container.PaymentRepository),
		),
	}
}

func (ctl *FindOneOrderController) Execute(ctx context.Context, orderId string) (dto.OrderDetails, error) {
	return ctl.usc.FindOne(ctx, orderId)
}
