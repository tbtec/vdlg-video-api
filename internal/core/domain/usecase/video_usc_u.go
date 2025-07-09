package usecase

import (
	"context"
	"log/slog"
	"strings"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type VideoUpdateUseCase struct {
	videoGtw *gateway.VideoGateway
}

func NewVideoUpdateUseCase(videoGtw *gateway.VideoGateway) *VideoUpdateUseCase {
	return &VideoUpdateUseCase{
		videoGtw: videoGtw,
	}
}

func (uc *VideoUpdateUseCase) Execute(ctx context.Context, updateVideo dto.UpdateVideo) (*entity.Video, error) {
	var video *entity.Video

	slog.InfoContext(ctx, "Executing video update use case", slog.Any("updateVideo", updateVideo))

	if updateVideo.InputMessage != nil {
		videoId := uc.getVideoId(updateVideo.InputMessage.Key)
		slog.InfoContext(ctx, "Updating video status", slog.Any("videoId", videoId))

		video, err := uc.videoGtw.FindOne(ctx, videoId)
		if video != nil {
			slog.InfoContext(ctx, "Video found", slog.Any("video", video))

			video.SetStatus(entity.VideoStatusProcessing)
			slog.InfoContext(ctx, "Setting video status to Processing", slog.Any("videoId", video.ID))

			uc.videoGtw.Update(ctx, video)
		}
		if err != nil {
			slog.ErrorContext(ctx, "Error file", slog.Any("error", err))
		}
	}
	if updateVideo.OutputMessage != nil {
		videoId := uc.getVideoId(updateVideo.OutputMessage.FileName)
		slog.InfoContext(ctx, "Updating video status", slog.Any("videoId", videoId))

		video, err := uc.videoGtw.FindOne(ctx, videoId)
		if video != nil {
			slog.InfoContext(ctx, "Video found", slog.Any("video", video))

			if updateVideo.OutputMessage.Status == string(entity.VideoStatusCompleted) {
				video.SetStatus(entity.VideoStatusCompleted)
			} else {
				video.SetStatus(entity.VideoStatusError)
			}

			slog.InfoContext(ctx, "Setting video status to Processing", slog.Any("videoId", video.ID))
			uc.videoGtw.Update(ctx, video)
		}
		if err != nil {
			slog.ErrorContext(ctx, "Error file", slog.Any("error", err))
		}
	}

	return video, nil
}

func (uc *VideoUpdateUseCase) getVideoId(fileName string) string {
	parts := strings.Split(fileName, ".")
	file := strings.Split(parts[0], "/")[1]

	return strings.Split(file, "_")[0]
}
