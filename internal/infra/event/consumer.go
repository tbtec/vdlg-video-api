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

type IConsumerService interface {
	ConsumeMessage(ctx context.Context) (*dto.Order, error)
}

type ConsumerService struct {
	QueueUrl string
	QueueArn string
	Client   *sqs.Client
}

func NewConsumerService(QueueUrl string, config aws.Config) IConsumerService {
	return &ConsumerService{
		QueueUrl: QueueUrl,
		Client:   sqs.NewFromConfig(config),
	}
}

func (consumer *ConsumerService) ConsumeMessage(ctx context.Context) (*dto.Order, error) {
	// Receive a message from the queue
	resp, err := consumer.Client.ReceiveMessage(ctx, &sqs.ReceiveMessageInput{
		QueueUrl:            &consumer.QueueUrl,
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

	var order dto.Order
	err = json.Unmarshal([]byte(envelope.Message), &order)
	if err != nil {
		return nil, fmt.Errorf("erro ao desserializar Order: %w", err)
	}
	slog.InfoContext(ctx, "Received message", "MessageId", *resp.Messages[0].MessageId)
	slog.InfoContext(ctx, "Received message", "body", order)

	// Delete the message from the queue
	out, delErr := consumer.Client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &consumer.QueueUrl,
		ReceiptHandle: resp.Messages[0].ReceiptHandle,
	})
	if delErr != nil {
		slog.ErrorContext(ctx, "Error deleting message", "error", delErr)
	}
	slog.InfoContext(ctx, "Message deleted", "recepit", *&out.ResultMetadata)

	return &order, nil
}

func (consumer *ConsumerService) DeleteMessage(ctx context.Context, receiptHandle string) error {
	_, err := consumer.Client.DeleteMessage(ctx, &sqs.DeleteMessageInput{
		QueueUrl:      &consumer.QueueUrl,
		ReceiptHandle: &receiptHandle,
	})
	if err != nil {
		return err
	}

	return nil
}
