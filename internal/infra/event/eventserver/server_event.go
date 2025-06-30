package eventserver

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/env"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/event"
)

type EventServer struct {
	ConsumerService              event.IConsumerService
	ConsumerProductionController *controller.ConsumerProductionController
}

func NewEventServer(container *container.Container, config env.Config) *EventServer {
	slog.InfoContext(context.Background(), "Creating Event Server...")

	cpc := controller.NewConsumerProductionController(container)
	cs := container.ConsumerService

	return &EventServer{
		ConsumerService:              cs,
		ConsumerProductionController: cpc}

}

func (eventServer *EventServer) Consume(ctx context.Context) {

	// Start the consumer service
	order, err := eventServer.ConsumerService.ConsumeMessage(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error reading message", slog.Any("error", err))
	}
	if order == nil {
		// slog.InfoContext(ctx, "No messages available")
	} else {
		// slog.InfoContext(ctx, "Received message: ", &order)

		err2 := eventServer.ConsumerProductionController.Execute(ctx, order.ID, order.Status)
		if err2 != nil {
			slog.ErrorContext(ctx, "Error processing message", slog.Any("error", err2))

		}
	}
}
