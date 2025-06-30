package usecase

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/types/xerrors"
)

var (
	ErrorOCRTCustomerNotFound = xerrors.NewBusinessError("TL-ORDERCRT-002", "Customer not found")
)

type UscCreateOrder struct {
	orderGateway    *gateway.OrderGateway
	customerGateway *gateway.CustomerGateway
	orderPresenter  *presenter.OrderPresenter
}

func NewUseCaseCreateOrder(orderGateway *gateway.OrderGateway,
	customerGateway *gateway.CustomerGateway) *UscCreateOrder {
	return &UscCreateOrder{
		orderGateway:    orderGateway,
		customerGateway: customerGateway,
		orderPresenter:  presenter.NewOrderPresenter(),
	}
}

func (usc *UscCreateOrder) Create(ctx context.Context, createOrder dto.CreateOrder) (dto.Order, error) {

	var customerPresent bool = false

	if createOrder.CustomerId != "" {
		customerPresent = true
	}

	if customerPresent {
		customer, err := usc.customerGateway.FindOne(ctx, createOrder.CustomerId)
		if err != nil {
			return dto.Order{}, err
		}
		if customer == nil {
			return dto.Order{}, ErrorOCRTCustomerNotFound
		}

	}

	order := entity.NewOrder(createOrder, customerPresent)

	err := usc.orderGateway.Create(ctx, &order)
	if err != nil {
		return dto.Order{}, err
	}

	return usc.orderPresenter.BuildOrderCreateResponse(order, nil), nil
}
