package container

import (
	"context"
	"log"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/tbtec/tremligeiro/internal/env"
	rdbms "github.com/tbtec/tremligeiro/internal/infra/database"
	"github.com/tbtec/tremligeiro/internal/infra/database/postgres"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
	"github.com/tbtec/tremligeiro/internal/infra/event"
	"github.com/tbtec/tremligeiro/internal/infra/file"
)

type Container struct {
	Config            env.Config
	TremLigeiroDB     rdbms.RDBMS
	ConsumerService   event.IConsumerService
	VideoRepository   repository.IVideoRepository
	FileUploadService file.IFileService
}

func New(config env.Config) (*Container, error) {
	factory := Container{}
	factory.Config = config

	return &factory, nil
}

func (container *Container) Start(ctx context.Context) error {

	var awsConfig aws.Config
	var err error

	if container.Config.Env == "local-stack" { // LocalStack
		awsConfig = container.GetLocalStackConfig(ctx)
	} else {
		awsConfig, err = config.LoadDefaultConfig(ctx,
			config.WithRegion(container.Config.AwsRegion))
		log.Printf("AWS Region: %s", container.Config.AwsRegion)
		if err != nil {
			log.Fatalf("erro ao carregar config: %v", err)
		}
	}

	err = postgres.Migrate(getPostgreSQLConf(container.Config))
	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
	}
	container.TremLigeiroDB, err = postgres.New(getPostgreSQLConf(container.Config))
	if err != nil {
		slog.ErrorContext(context.Background(), err.Error())
		return err
	}

	container.ConsumerService = event.NewConsumerService(container.Config.InputQueueUrl, container.Config.OutputQueueUrl, awsConfig)
	container.VideoRepository = repository.NewVideoRepository(container.TremLigeiroDB)
	container.FileUploadService = file.NewFileService(container.Config.S3BucketName, awsConfig)

	return nil
}

func (container *Container) Stop() error {
	db, err := container.TremLigeiroDB.DB.DB()
	if err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func getPostgreSQLConf(config env.Config) postgres.PostgreSQLConf {
	return postgres.PostgreSQLConf{
		User:   config.DbUser,
		Pass:   config.DbPassword,
		Url:    config.DbHost,
		Port:   config.DbPort,
		DbName: config.DbName,
	}
}

func (container *Container) GetLocalStackConfig(ctx context.Context) aws.Config {

	awsConfig, err := config.LoadDefaultConfig(ctx,
		config.WithRegion("us-east-1"),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider("test", "test", "")),
	)
	awsConfig.BaseEndpoint = aws.String("http://localhost:4566")

	if err != nil {
		log.Fatalf("erro ao carregar config: %v", err)
	}

	return awsConfig
}
