package file

import (
	"context"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type IFileUploadService interface {
	GenerateUploadUrl(ctx context.Context, customerId string, fileName string) (string, error)
}

type FileUploadService struct {
	BucketName string
	Client     *s3.Client
}

func NewFileUploadService(bucketName string, config aws.Config) IFileUploadService {
	return &FileUploadService{
		BucketName: bucketName,
		Client:     s3.NewFromConfig(config),
	}
}

func (service *FileUploadService) GenerateUploadUrl(ctx context.Context, customerId string, fileName string) (string, error) {
	slog.InfoContext(ctx, "Generating upload URL", slog.String("customerId", customerId), slog.String("fileName", fileName))
	// Implement the file upload logic here
	// For example, using the S3 client to upload the file
	// and returning the file URL or key

	// Placeholder return statement
	return "https://example.com/" + fileName, nil
}
