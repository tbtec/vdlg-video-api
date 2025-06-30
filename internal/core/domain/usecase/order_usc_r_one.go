package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UscFindOneOrder struct {
	orderProductGateway *gateway.OrderProductGateway
	orderGateway        *gateway.OrderGateway
	orderPresenter      *presenter.OrderPresenter
}

func NewUscFindOneOrder(orderGateway *gateway.OrderGateway,
	orderProductGateway *gateway.OrderProductGateway,
	paymentGateway *gateway.PaymentGateway) *UscFindOneOrder {
	return &UscFindOneOrder{
		orderProductGateway: orderProductGateway,
		orderGateway:        orderGateway,
		orderPresenter:      presenter.NewOrderPresenter(),
	}
}

func (usc *UscFindOneOrder) FindOne(ctx context.Context, orderId string) (dto.OrderDetails, error) {
	order, err := usc.getOrder(ctx, orderId)
	if err != nil {
		return dto.OrderDetails{}, err
	}
	orderProducts, err := usc.orderProductGateway.FindByOrderId(ctx, orderId)
	if err != nil {
		return dto.OrderDetails{}, err
	}

	return usc.orderPresenter.BuildOrderDetailsCreateResponse(*order, orderProducts), nil
}

func (usc *UscFindOneOrder) getOrder(ctx context.Context, orderId string) (*entity.Order, error) {
	order, err := usc.orderGateway.FindOne(ctx, orderId)
	if order == nil {
		return nil, ErrorOrderNotFound
	}
	if err != nil {
		return nil, err
	}

	return order, nil
}
