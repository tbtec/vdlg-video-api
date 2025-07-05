package controller

import (
	"context"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"

	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type VideoFindOneRestController struct {
	controller *ctl.VideoFindOneController
}

func NewVideoFindOneRestController(container *container.Container) httpserver.IController {
	return &VideoFindOneRestController{
		controller: ctl.NewVideoFindOneController(container),
	}
}

func (ctl *VideoFindOneRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {

	videoId := request.ParseParamString("id")

	output, err := ctl.controller.Execute(ctx, videoId)
	if err != nil {
		return httpserver.HandleError(ctx, err)
	}

	return httpserver.Ok(output)
}
