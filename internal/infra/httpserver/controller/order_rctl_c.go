package controller

import (
	"context"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type OrderCreateRestController struct {
	controller *ctl.CreateOrderController
}

func NewOrderCreateRestController(container *container.Container) httpserver.IController {
	return &OrderCreateRestController{
		controller: ctl.NewCreateOrderController(container),
	}
}

func (ctl *OrderCreateRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {
	orderRequest := dto.CreateOrder{}

	err := request.ParseBody(ctx, &orderRequest)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	order, err := ctl.controller.Execute(ctx, orderRequest)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(order)
}
