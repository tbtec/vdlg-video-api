package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type VideoUpdateController struct {
	usc *usecase.VideoUpdateUseCase
}

func NewVideoUpdateController(container *container.Container) *VideoUpdateController {
	return &VideoUpdateController{
		usc: usecase.NewVideoUpdateUseCase(
			gateway.NewVideoGateway(container.VideoRepository),
		),
	}
}

func (ctl *VideoUpdateController) Execute(ctx context.Context, updateVideo dto.UpdateVideo) error {
	_, err := ctl.usc.Execute(ctx, updateVideo)
	if err != nil {
		return err
	}
	return nil
}
