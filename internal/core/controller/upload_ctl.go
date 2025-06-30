package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type UploadController struct {
	usc *usecase.UploadUseCase
}

func NewUploadController(container *container.Container) *UploadController {
	return &UploadController{
		usc: usecase.NewUploadUseCase(
			gateway.NewUploadGateway(container.FileUploadService),
		),
	}
}

func (ctl *UploadController) Execute(ctx context.Context, createVideo dto.CreateVideo) (string, error) {
	return ctl.usc.Execute(ctx, createVideo)
}
