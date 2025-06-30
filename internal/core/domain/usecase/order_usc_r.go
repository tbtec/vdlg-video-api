package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type UscFindOrder struct {
	orderProductGateway *gateway.OrderProductGateway
	orderGateway        *gateway.OrderGateway
	orderPresenter      *presenter.OrderPresenter
}

func NewUseCaseFindOrder(orderGateway *gateway.OrderGateway,
	orderProductGateway *gateway.OrderProductGateway) *UscFindOrder {
	return &UscFindOrder{
		orderProductGateway: orderProductGateway,
		orderGateway:        orderGateway,
		orderPresenter:      presenter.NewOrderPresenter(),
	}
}

func (usc *UscFindOrder) Find(ctx context.Context) (dto.OrderContent, error) {
	orders, err := usc.orderGateway.Find(ctx)
	if err != nil {
		return dto.OrderContent{}, err
	}

	return usc.orderPresenter.BuildOrderContentResponse(orders), nil
}
