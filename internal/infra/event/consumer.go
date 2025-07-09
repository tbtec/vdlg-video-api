package event

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/sqs"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type SNSEnvelope struct {
	Type      string `json:"Type"`
	MessageId string `json:"MessageId"`
	TopicArn  string `json:"TopicArn"`
	Message   string `json:"Message"`
}

type S3Event struct {
	Records []struct {
		S3 struct {
			Object struct {
				Key string `json:"key"`
			} `json:"object"`
		} `json:"s3"`
	} `json:"Records"`
}

type IConsumerService interface {
	ConsumeMessageInput(ctx context.Context) (*dto.InputMessage, error)
	ConsumeMessageOutput(ctx context.Context) (*dto.OutputMessage, error)
}

type ConsumerService struct {
	QueueInputUrl  string
	QueueOutputUrl string
	Client         *sqs.Client
}

func NewConsumerService(queueInputUrl string, queueOutputUrl string, config aws.Config) IConsumerService {
	return &ConsumerService{
		QueueInputUrl:  queueInputUrl,
		QueueOutputUrl: queueOutputUrl,
		Client:         sqs.NewFromConfig(config),
	}
}

func (consumer *ConsumerService) ConsumeMessageInput(ctx context.Context) (*dto.InputMessage, error) {
	// Receive a message from the queue
	resp, err := consumer.Client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &consumer.QueueInputUrl,
		MaxNumberOfMessages: 1,
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Messages) == 0 {
		return nil, nil // No messages available
	}

	var envelope SNSEnvelope
	err = json.Unmarshal([]byte(*resp.Messages[0].Body), &envelope)
	if err != nil {
		return nil, fmt.Errorf("erro ao desserializar envelope SNS: %w", err)
	}

	slog.InfoContext(ctx, "Received message", "MessageId", envelope.Message)

	var event S3Event
	err = json.Unmarshal([]byte(envelope.Message), &event)
	if err != nil {
		return nil, fmt.Errorf("erro ao desserializar Order: %w", err)
	}
	slog.InfoContext(ctx, "Received message", "MessageId", *resp.Messages[0].MessageId)
	slog.InfoContext(ctx, "Received message", "body", event)

	// Delete the message from the queue
	out, delErr := consumer.Client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &consumer.QueueInputUrl,
		ReceiptHandle: resp.Messages[0].ReceiptHandle,
	})
	if delErr != nil {
		slog.ErrorContext(ctx, "Error deleting message", "error", delErr)
	}
	slog.InfoContext(ctx, "Message deleted", "recepit", *&out.ResultMetadata)

	return &dto.InputMessage{
		Key: event.Records[0].S3.Object.Key,
	}, nil
}

func (consumer *ConsumerService) ConsumeMessageOutput(ctx context.Context) (*dto.OutputMessage, error) {
	// Receive a message from the queue
	resp, err := consumer.Client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &consumer.QueueOutputUrl,
		MaxNumberOfMessages: 1,
	})
	if err != nil {
		return nil, err
	}

	if len(resp.Messages) == 0 {
		return nil, nil // No messages available
	}

	var envelope SNSEnvelope
	err = json.Unmarshal([]byte(*resp.Messages[0].Body), &envelope)
	if err != nil {
		return nil, fmt.Errorf("erro ao desserializar envelope SNS: %w", err)
	}

	var message dto.OutputMessage
	err = json.Unmarshal([]byte(envelope.Message), &message)
	if err != nil {
		return nil, fmt.Errorf("erro ao desserializar Order: %w", err)
	}
	slog.InfoContext(ctx, "Received message Output", "MessageId", *resp.Messages[0].MessageId)
	slog.InfoContext(ctx, "Received message Output", "body", message)

	// Delete the message from the queue
	out, delErr := consumer.Client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &consumer.QueueOutputUrl,
		ReceiptHandle: resp.Messages[0].ReceiptHandle,
	})
	if delErr != nil {
		slog.ErrorContext(ctx, "Error deleting message", "error", delErr)
	}
	slog.InfoContext(ctx, "Message deleted", "recepit", *&out.ResultMetadata)

	return &message, nil
}

func (consumer *ConsumerService) DeleteMessage(ctx context.Context, receiptHandle string) error {
	_, err := consumer.Client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &consumer.QueueInputUrl,
		ReceiptHandle: &receiptHandle,
	})
	if err != nil {
		return err
	}

	return nil
}
