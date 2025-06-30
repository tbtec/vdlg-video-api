package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type FindOrderController struct {
	usc *usecase.UscFindOrder
}

func NewFindOrderController(container *container.Container) *FindOrderController {
	return &FindOrderController{
		usc: usecase.NewUseCaseFindOrder(
			gateway.NewOrderGateway(container.OrderRepository),
			gateway.NewOrderProductGateway(container.OrderProductRepository),
		),
	}
}

func (ctl *FindOrderController) Execute(ctx context.Context) (dto.OrderContent, error) {
	return ctl.usc.Find(ctx)
}
