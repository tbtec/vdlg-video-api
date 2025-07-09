package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type UploadController struct {
	usc *usecase.UploadUseCase
	prt *presenter.VideoPresenter
}

func NewUploadController(container *container.Container) *UploadController {
	return &UploadController{
		usc: usecase.NewUploadUseCase(
			gateway.NewFileGateway(container.FileUploadService),
			gateway.NewVideoGateway(container.VideoRepository),
		),
		prt: presenter.NewVideoPresenter(),
	}
}

func (ctl *UploadController) Execute(ctx context.Context, createVideo dto.CreateVideo) (dto.Video, error) {
	url, err := ctl.usc.Execute(ctx, createVideo)
	if err != nil {
		return dto.Video{}, err
	}
	return ctl.prt.BuildVideoResponse(url), nil
}
