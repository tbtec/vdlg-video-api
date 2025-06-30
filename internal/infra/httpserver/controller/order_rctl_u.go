package controller

import (
	"context"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/validator"

	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type UpdateOrderRestController struct {
	controller *ctl.UpdateOrderController
}

func NewUpdateOrderRestController(container *container.Container) httpserver.IController {
	return &UpdateOrderRestController{
		controller: ctl.NewUpdateOrderController(container),
	}
}

func (ctl *UpdateOrderRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	input := dto.UpdateOrder{}
	err := request.ParseBody(ctx, &input)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	err2 := validator.Validate(input)
	if err2 != nil {
		return httpserver.HandleError(ctx, err2)
	}

	orderId := request.Params["orderId"]

	err3 := ctl.controller.Execute(ctx, orderId, input.Status)
	if err3 != nil {
		return httpserver.HandleError(ctx, err3)
	}

	return httpserver.NoContent()
}
