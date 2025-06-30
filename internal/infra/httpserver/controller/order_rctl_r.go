package controller

import (
	"context"
	"time"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"

	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type OrderFindController struct {
	controller *ctl.FindOrderController
}

type OrderFindResponse struct {
	Content []OrderFindOrderResponse `json:"content"`
}

type OrderFindOrderResponse struct {
	ID          string    `json:"id"`
	CustomerId  *string   `json:"customer_id"`
	Status      string    `json:"status"`
	TotalAmount float64   `json:"total_amount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func NewOrderFindController(container *container.Container) httpserver.IController {
	return &OrderFindController{
		controller: ctl.NewFindOrderController(container),
	}
}

func (ctl *OrderFindController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	output, err := ctl.controller.Execute(ctx)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(output)
}
