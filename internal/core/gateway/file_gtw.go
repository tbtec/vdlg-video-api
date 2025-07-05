package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/infra/file"
)

type FileGateway struct {
	fileService file.IFileService
}

func NewFileGateway(fileService file.IFileService) *FileGateway {
	return &FileGateway{
		fileService: fileService,
	}
}

func (gtw *FileGateway) GenerateUploadUrl(ctx context.Context, fileName string) (string, error) {
	url, err := gtw.fileService.GenerateUploadUrl(ctx, fileName)
	if err != nil {
		return "", err
	}

	return url, nil
}

func (gtw *FileGateway) GenerateDownloadUrl(ctx context.Context, fileName string) (string, error) {
	url, err := gtw.fileService.GenerateDownloadUrl(ctx, fileName)
	if err != nil {
		return "", err
	}

	return url, nil
}
