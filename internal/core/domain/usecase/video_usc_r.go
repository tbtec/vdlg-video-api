package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
)

type VideoFindUseCase struct {
	videoGtw *gateway.VideoGateway
}

func NewVideoFindUseCase(videoGtw *gateway.VideoGateway) *VideoFindUseCase {
	return &VideoFindUseCase{
		videoGtw: videoGtw,
	}
}

func (usc *VideoFindUseCase) Execute(ctx context.Context, customerId string) ([]entity.Video, error) {
	videos, err := usc.videoGtw.Find(ctx, customerId)
	if err != nil {
		return []entity.Video{}, err
	}

	return videos, nil
}
