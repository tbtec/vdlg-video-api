package controller

import (
	"context"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"

	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type OrderFindOneRestController struct {
	controller *ctl.FindOneOrderController
}

func NewOrderFindOneRestController(container *container.Container) httpserver.IController {
	return &OrderFindOneRestController{
		controller: ctl.NewFindOneOrderController(container),
	}
}

func (ctl *OrderFindOneRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	orderId := request.Params["orderId"]

	output, err := ctl.controller.Execute(ctx, orderId)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(output)
}
