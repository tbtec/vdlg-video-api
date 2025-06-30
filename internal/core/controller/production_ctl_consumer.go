package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type ConsumerProductionController struct {
	usc *usecase.UscUpdateOrder
}

func NewConsumerProductionController(container *container.Container) *ConsumerProductionController {
	return &ConsumerProductionController{
		usc: usecase.NewUscUpdateOrder(
			gateway.NewOrderGateway(container.OrderRepository),
		),
	}
}

func (ctl *ConsumerProductionController) Execute(ctx context.Context, orderId string, newStatus string) error {
	return ctl.usc.Update(ctx, orderId, newStatus)
}
