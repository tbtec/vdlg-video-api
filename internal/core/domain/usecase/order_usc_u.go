package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/types/xerrors"
)

type UscUpdateOrder struct {
	orderGateway *gateway.OrderGateway
}

func NewUscUpdateOrder(orderGateway *gateway.OrderGateway) *UscUpdateOrder {
	return &UscUpdateOrder{
		orderGateway: orderGateway,
	}
}

func (usc *UscUpdateOrder) Update(ctx context.Context, orderId string, newStatus string) error {
	slog.InfoContext(ctx, "Updating order status", "orderId", orderId, "newStatus", newStatus)
	order, err := usc.getOrder(ctx, orderId)
	if err != nil {
		return err
	}
	if !order.ValidateStatus(order.Status, entity.OrderStatus(newStatus)) {
		return xerrors.NewBusinessError(ErrorStatusNotAllowedCode, fmt.Sprintf("Order status is not %s current status is %s", entity.OrderStatus(newStatus), order.Status))
	}
	order.Status = entity.OrderStatus(newStatus)

	err = usc.orderGateway.Update(ctx, order)
	if err != nil {
		return err
	}

	return nil
}

func (usc *UscUpdateOrder) getOrder(ctx context.Context, orderId string) (*entity.Order, error) {
	order, err := usc.orderGateway.FindOne(ctx, orderId)
	if order == nil {
		return nil, ErrorOrderNotFound
	}
	if err != nil {
		return nil, err
	}

	return order, nil
}
