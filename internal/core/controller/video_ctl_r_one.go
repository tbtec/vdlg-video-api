package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type VideoFindOneController struct {
	usc *usecase.VideoFindOneUseCase
	prt *presenter.VideoPresenter
}

func NewVideoFindOneController(container *container.Container) *VideoFindOneController {
	return &VideoFindOneController{
		usc: usecase.
			NewVideoFindOneUseCase(
				gateway.NewFileGateway(container.FileUploadService),
				gateway.NewVideoGateway(container.VideoRepository),
			),
		prt: presenter.NewVideoPresenter(),
	}
}

func (ctl *VideoFindOneController) Execute(ctx context.Context, id string) (dto.Video, error) {
	video, err := ctl.usc.Execute(ctx, id)
	if err != nil {
		return dto.Video{}, err
	}
	return ctl.prt.BuildVideoResponse(*video), nil
}
