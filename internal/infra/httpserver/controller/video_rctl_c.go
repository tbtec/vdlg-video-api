package controller

import (
	"context"
	"log/slog"

	ctl "github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/httpserver"
)

type UploadRestController struct {
	controller *ctl.UploadController
}

func NewUploadRestController(container *container.Container) httpserver.IController {
	return &UploadRestController{
		controller: ctl.NewUploadController(container),
	}
}

func (ctl *UploadRestController) Handle(ctx context.Context, request httpserver.Request) httpserver.Response {
	videoRequest := dto.CreateVideo{
		CustomerId: request.Headers["customer-id"],
	}
	slog.InfoContext(ctx, "Request received for video upload", slog.Any("videoRequest", request))
	slog.InfoContext(ctx, "Request received for video upload", slog.Any("videoRequest", videoRequest))

	err := request.ParseBody(ctx, &videoRequest)
	if err != nil {
		slog.ErrorContext(ctx, "Error parsing request body", slog.Any("error", err))
		return httpserver.HandleError(ctx, err)
	}
	slog.InfoContext(ctx, "Parsed video request", slog.Any("videoRequest", videoRequest))

	response, err2 := ctl.controller.Execute(ctx, videoRequest)
	if err2 != nil {
		return httpserver.HandleError(ctx, err2)
	}

	return httpserver.Ok(response)
}
