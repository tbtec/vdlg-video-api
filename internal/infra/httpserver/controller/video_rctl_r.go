package controller

import (
	"context"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"

	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type VideoFindRestController struct {
	controller *ctl.VideoFindController
}

func NewVideoFindController(container *container.Container) httpserver.IController {
	return &VideoFindRestController{
		controller: ctl.NewVideoFindController(container),
	}
}

func (ctl *VideoFindRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	customerId := request.Query["customerId"]

	output, err := ctl.controller.Execute(ctx, customerId)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(output)
}
