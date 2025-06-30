package usecase

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/types/xerrors"
)

var (
	ErrPaymentUnauthorized    = xerrors.NewBusinessError("TL-ORDERCKT-001", "Payment unauthorized")
	ErrorCustomerNotFound     = xerrors.NewBusinessError("TL-ORDERCKT-002", "Customer not found")
	ErrorOrderNotFound        = xerrors.NewBusinessError("TL-ORDERCKT-003", "Order not found")
	ErrorProductNotFound      = xerrors.NewBusinessError("TL-ORDERCKT-004", "Product not found")
	ErrorStatusNotAllowedCode = "TL-ORDERCKT-005"
)

type UscOrderCheckOut struct {
	orderProductGateway  *gateway.OrderProductGateway
	orderGateway         *gateway.OrderGateway
	productGateway       *gateway.ProductGateway
	paymentGateway       *gateway.PaymentGateway
	orderPresenter       *presenter.OrderPresenter
	orderProducerGateway *gateway.OrderProducerGateway
}

func NewUseCaseOrderCheckout(orderGateway *gateway.OrderGateway,
	productGateway *gateway.ProductGateway,
	orderProductGateway *gateway.OrderProductGateway,
	paymentGateway *gateway.PaymentGateway,
	orderPresenter *presenter.OrderPresenter,
	orderProducerGateway *gateway.OrderProducerGateway,
) *UscOrderCheckOut {
	return &UscOrderCheckOut{
		orderProductGateway:  orderProductGateway,
		orderGateway:         orderGateway,
		productGateway:       productGateway,
		paymentGateway:       paymentGateway,
		orderPresenter:       orderPresenter,
		orderProducerGateway: orderProducerGateway,
	}
}

func (usc *UscOrderCheckOut) Checkout(ctx context.Context, orderCheckout dto.OrderCheckout) (dto.Order, error) {
	slog.InfoContext(ctx, "Intiating checkout order", "orderId", orderCheckout.OrderId)
	var err error

	order, err := usc.getOrder(ctx, orderCheckout.OrderId)
	if err != nil {
		return dto.Order{}, err
	}
	if !order.ValidateStatus(order.Status, entity.OrderStatusReceived) {
		return dto.Order{}, xerrors.NewBusinessError(ErrorStatusNotAllowedCode, fmt.Sprintf("Order status is not %s current status is %s", entity.OrderStatusPending, order.Status))
	}

	orderProducts, totalAmount, err := usc.getOrderProducts(ctx, orderCheckout.OrderId, orderCheckout.Products)
	if err != nil {
		return dto.Order{}, err
	}
	order.SetTotalAmount(totalAmount)

	payment := entity.NewPayment(order.ID)

	for _, orderProduct := range orderProducts {
		usc.orderProductGateway.Create(ctx, orderProduct)
	}

	err2 := usc.doPayment(ctx, *order, orderProducts, orderCheckout.MetaData)
	if err2 != nil {
		return dto.Order{}, err2
	}

	usc.orderGateway.Update(ctx, order)

	_ = usc.orderProducerGateway.PublishMessage(ctx, order)

	return usc.orderPresenter.BuildOrderCreateResponse(*order, &payment.ID), nil
}

func (usc *UscOrderCheckOut) doPayment(ctx context.Context, order entity.Order, orderProducts []entity.OrderProduct, metadata dto.MetaData) error {

	err := usc.paymentGateway.RequestPayment(ctx, order, orderProducts, metadata)
	if err != nil {
		return err
	}

	return nil
}

func (usc *UscOrderCheckOut) getOrder(ctx context.Context, orderId string) (*entity.Order, error) {
	order, err := usc.orderGateway.FindOne(ctx, orderId)
	if order == nil {
		return nil, ErrorOrderNotFound
	}
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (usc *UscOrderCheckOut) getOrderProducts(ctx context.Context, orderId string, products []dto.OrderCheckoutProduct) ([]entity.OrderProduct, float64, error) {
	orderProducts := []entity.OrderProduct{}
	var totalAmount float64

	for _, cmdProduct := range products {
		productEntity, err := usc.productGateway.FindOne(ctx, cmdProduct.ProductId)
		if productEntity == nil {
			return orderProducts, 0, ErrorProductNotFound
		}
		if err != nil {
			return orderProducts, 0, err
		}
		orderProduct := entity.NewOrderProduct(orderId, productEntity.ID, cmdProduct.Quantity, productEntity.Amount)

		orderProducts = append(orderProducts, orderProduct)
		quantity := float64(cmdProduct.Quantity) * productEntity.Amount
		totalAmount += quantity
	}

	return orderProducts, totalAmount, nil
}
