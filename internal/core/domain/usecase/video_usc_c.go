package usecase

import (
	"context"
	"log/slog"
	"strings"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UploadUseCase struct {
	fileGtw  *gateway.FileGateway
	videoGtw *gateway.VideoGateway
}

func NewUploadUseCase(fileGtw *gateway.FileGateway, videoGtw *gateway.VideoGateway) *UploadUseCase {
	return &UploadUseCase{
		fileGtw:  fileGtw,
		videoGtw: videoGtw,
	}
}

func (uc *UploadUseCase) Execute(ctx context.Context, createVideo dto.CreateVideo) (entity.Video, error) {
	slog.InfoContext(ctx, "Generating upload URL", slog.Any("customerId", createVideo.CustomerId), slog.Any("fileName", createVideo.FileName))

	video := entity.NewVideo(createVideo.CustomerId)

	fileName := uc.formatFileNameUpload(video.ID, createVideo.FileName)

	url, err := uc.fileGtw.GenerateUploadUrl(ctx, fileName)

	video.SetFileNameInput(fileName)
	video.SetUploadUrl(url)
	video.SetFileNameOutput(uc.formatFileNameDownload(fileName))

	uc.videoGtw.Create(ctx, &video)

	if err != nil {
		slog.ErrorContext(ctx, "Error uploading file", slog.Any("error", err))
	}
	return video, nil
}

func (uc *UploadUseCase) formatFileNameUpload(id string, fileName string) string {
	// parts := strings.Split(fileName, ".")
	// formatted := time.Now().Format("20060102_150405")
	// return id + "_" + formatted + "_" + parts[0] + "." + parts[1]
	parts := strings.Split(fileName, ".")
	return id + "." + parts[1]
}

func (uc *UploadUseCase) formatFileNameDownload(fileName string) string {
	parts := strings.Split(fileName, ".")
	return parts[0] + ".zip"
}
