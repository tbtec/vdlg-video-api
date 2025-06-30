package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type UpdateOrderController struct {
	usc *usecase.UscUpdateOrder
}

func NewUpdateOrderController(container *container.Container) *UpdateOrderController {
	return &UpdateOrderController{
		usc: usecase.NewUscUpdateOrder(
			gateway.NewOrderGateway(container.OrderRepository),
		),
	}
}

func (ctl *UpdateOrderController) Execute(ctx context.Context, orderId string, newStatus string) error {
	return ctl.usc.Update(ctx, orderId, newStatus)
}
