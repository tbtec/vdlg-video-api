package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type VideoFindController struct {
	usc *usecase.VideoFindUseCase
	prt *presenter.VideoPresenter
}

func NewVideoFindController(container *container.Container) *VideoFindController {
	return &VideoFindController{
		usc: usecase.NewVideoFindUseCase(
			gateway.NewVideoGateway(container.VideoRepository),
		),
		prt: presenter.NewVideoPresenter(),
	}
}

func (ctl *VideoFindController) Execute(ctx context.Context) (dto.VideoContent, error) {
	videos, err := ctl.usc.Execute(ctx)
	if err != nil {
		return dto.VideoContent{}, err
	}
	return ctl.prt.BuildVideoContentResponse(videos), nil
}
