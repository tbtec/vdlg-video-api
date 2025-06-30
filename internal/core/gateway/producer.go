package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/event"
)

type OrderProducerGateway struct {
	producerService event.IProducerService
}

func NewOrderProducerGateway(producerService event.IProducerService) *OrderProducerGateway {
	return &OrderProducerGateway{
		producerService: producerService,
	}
}

func (gtw *OrderProducerGateway) PublishMessage(ctx context.Context, order *entity.Order) error {
	orderDto := dto.Order{
		ID:          order.ID,
		CustomerId:  order.CustomerId,
		Status:      string(order.Status),
		TotalAmount: order.TotalAmount,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}

	err := gtw.producerService.PublishMessage(ctx, orderDto)
	if err != nil {
		return err
	}

	return nil
}
