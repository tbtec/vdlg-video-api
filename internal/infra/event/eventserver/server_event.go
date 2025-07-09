package eventserver

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/controller"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/env"
	"github.com/tbtec/tremligeiro/internal/infra/container"
	"github.com/tbtec/tremligeiro/internal/infra/event"
)

type EventServer struct {
	ConsumerService       event.IConsumerService
	VideoUpdateController *controller.VideoUpdateController
}

func NewEventServer(container *container.Container, config env.Config) *EventServer {
	slog.InfoContext(context.Background(), "Creating Event Server...")

	cpc := controller.NewVideoUpdateController(container)
	cs := container.ConsumerService

	return &EventServer{
		ConsumerService:       cs,
		VideoUpdateController: cpc}

}

func (eventServer *EventServer) ConsumeInput(ctx context.Context) {

	inputMessage, err := eventServer.ConsumerService.ConsumeMessageInput(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "Error reading message", slog.Any("error", err))
	}
	if inputMessage == nil {
		// slog.InfoContext(ctx, "No messages available")
	} else {
		slog.InfoContext(ctx, "Received message: ", &inputMessage)

		err2 := eventServer.VideoUpdateController.Execute(ctx, dto.UpdateVideo{
			InputMessage: inputMessage,
		})
		if err2 != nil {
			slog.ErrorContext(ctx, "Error processing message", slog.Any("error", err2))

		}
	}
}
