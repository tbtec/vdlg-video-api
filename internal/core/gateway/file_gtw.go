package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/infra/file"
)

type FileGateway struct {
	fileService file.IFileUploadService
}

func NewUploadGateway(fileService file.IFileUploadService) *FileGateway {
	return &FileGateway{
		fileService: fileService,
	}
}

func (gtw *FileGateway) GenerateUploadUrl(ctx context.Context, customerId string, fileName string) (string, error) {

	return gtw.fileService.GenerateUploadUrl(ctx, customerId, fileName)
}
