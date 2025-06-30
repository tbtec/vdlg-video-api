package usecase

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UploadUseCase struct {
	fileGateway *gateway.FileGateway
}

func NewUploadUseCase(uploadGateway *gateway.FileGateway) *UploadUseCase {
	return &UploadUseCase{
		fileGateway: uploadGateway,
	}
}

func (uc *UploadUseCase) Execute(ctx context.Context, createVideo dto.CreateVideo) (string, error) {
	slog.InfoContext(ctx, "Generating upload URL", slog.Any("customerId", createVideo.CustomerId), slog.Any("fileName", createVideo.FileName))
	url, err := uc.fileGateway.GenerateUploadUrl(ctx, createVideo.CustomerId, createVideo.FileName)
	if err != nil {
		slog.ErrorContext(ctx, "Error uploading file", slog.Any("error", err))
	}
	return url, nil
}
