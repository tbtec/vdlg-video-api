package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/types/xerrors"
)

var (
	ErrorOrderNotFound = xerrors.NewBusinessError("VL-VIDEO-001", "Video not found")
)

type VideoFindOneUseCase struct {
	fileGtw  *gateway.FileGateway
	videoGtw *gateway.VideoGateway
}

func NewVideoFindOneUseCase(fileGtw *gateway.FileGateway,
	videoGtw *gateway.VideoGateway) *VideoFindOneUseCase {
	return &VideoFindOneUseCase{
		videoGtw: videoGtw,
		fileGtw:  fileGtw,
	}
}

func (usc *VideoFindOneUseCase) Execute(ctx context.Context, id string) (*entity.Video, error) {
	video, err := usc.videoGtw.FindOne(ctx, id)
	if video == nil {
		return nil, ErrorOrderNotFound
	}
	if err != nil {
		return nil, err
	}

	url, err2 := usc.fileGtw.GenerateDownloadUrl(ctx, video.FileNameOutput)
	if err2 != nil {
		return nil, err2
	}
	video.SetDownloadUrl(url)

	return video, nil
}
