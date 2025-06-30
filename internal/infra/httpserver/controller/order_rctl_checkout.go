package controller

import (
	"context"
	"log/slog"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
	"github.com/tbtec/tremligeiro/internal/validator"
)

type OrderCheckoutRestController struct {
	controller *ctl.OrderCheckoutController
}

func NewOrderCheckoutRestController(container *container.Container) httpserver.IController {
	return &OrderCheckoutRestController{
		controller: ctl.NewOrderCheckoutController(container),
	}
}

func (ctl *OrderCheckoutRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {
	orderCheckoutRequest := dto.OrderCheckout{
		OrderId: request.ParseParamString("orderId"),
	}

	errBody := request.ParseBody(ctx, &orderCheckoutRequest)
	if errBody != nil {
		return httpserver.HandleError(ctx, errBody)
	}

	err := validator.Validate(orderCheckoutRequest)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	output, err := ctl.controller.Execute(ctx, orderCheckoutRequest)
	if err != nil {
		slog.ErrorContext(ctx, "Error on checkout order", "error", err)
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(output)
}
