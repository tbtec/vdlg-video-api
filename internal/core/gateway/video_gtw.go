package gateway

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
)

type VideoGateway struct {
	videoRepository repository.IVideoRepository
}

func NewVideoGateway(videoRepository repository.IVideoRepository) *VideoGateway {
	return &VideoGateway{
		videoRepository: videoRepository,
	}
}

func (gtw *VideoGateway) Create(ctx context.Context, video *entity.Video) error {

	videoModel := model.Video{
		ID:             video.ID,
		CustomerId:     video.CustomerId,
		Status:         string(video.Status),
		FileNameInput:  video.FileNameInput,
		FileNameOutput: video.FileNameOutput,
		CreatedAt:      video.CreatedAt,
		UpdatedAt:      video.UpdatedAt,
	}

	err := gtw.videoRepository.Create(ctx, &videoModel)

	if err != nil {
		return err
	}

	return nil
}

func (gtw *VideoGateway) Update(ctx context.Context, video *entity.Video) error {

	videoModel := model.Video{
		ID:             video.ID,
		CustomerId:     video.CustomerId,
		Status:         string(video.Status),
		FileNameInput:  video.FileNameInput,
		FileNameOutput: video.FileNameOutput,
		CreatedAt:      video.CreatedAt,
		UpdatedAt:      video.UpdatedAt,
	}

	err := gtw.videoRepository.Update(ctx, &videoModel)

	if err != nil {
		slog.ErrorContext(ctx, "Error updating video", slog.Any("error", err), slog.Any("video", videoModel))
		return err
	}

	return nil
}

func (gtw *VideoGateway) FindOne(ctx context.Context, id string) (*entity.Video, error) {
	videoModel, err := gtw.videoRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	video := entity.Video{
		ID:             videoModel.ID,
		CustomerId:     videoModel.CustomerId,
		Status:         entity.VideoStatus(videoModel.Status),
		FileNameInput:  videoModel.FileNameInput,
		FileNameOutput: videoModel.FileNameOutput,
		CreatedAt:      videoModel.CreatedAt,
		UpdatedAt:      videoModel.UpdatedAt,
	}

	return &video, nil
}

func (gtw *VideoGateway) Find(ctx context.Context) ([]entity.Video, error) {
	videosModel, err := gtw.videoRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	videos := []entity.Video{}
	for _, videoModel := range videosModel {
		video := entity.Video{
			ID:             videoModel.ID,
			CustomerId:     videoModel.CustomerId,
			Status:         entity.VideoStatus(videoModel.Status),
			FileNameInput:  videoModel.FileNameInput,
			FileNameOutput: videoModel.FileNameOutput,
			CreatedAt:      videoModel.CreatedAt,
			UpdatedAt:      videoModel.UpdatedAt,
		}
		videos = append(videos, video)
	}

	return videos, nil
}
