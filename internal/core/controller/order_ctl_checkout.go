package controller

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/usecase"
	"github.com/tbtec/tremligeiro/internal/core/gateway"
	"github.com/tbtec/tremligeiro/internal/core/presenter"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/container"
)

type OrderCheckoutController struct {
	usc *usecase.UscOrderCheckOut
}

func NewOrderCheckoutController(container *container.Container) *OrderCheckoutController {
	return &OrderCheckoutController{
		usc: usecase.NewUseCaseOrderCheckout(
			gateway.NewOrderGateway(container.OrderRepository),
			gateway.NewProductGateway(container.ProductRepository, container.ProductService),
			gateway.NewOrderProductGateway(container.OrderProductRepository),
			gateway.NewPaymentGateway(container.PaymentService,
				container.PaymentRepository),
			presenter.NewOrderPresenter(),
			gateway.NewOrderProducerGateway(container.ProducerService),
		),
	}
}

func (ctl *OrderCheckoutController) Execute(ctx context.Context, input dto.OrderCheckout) (dto.Order, error) {
	return ctl.usc.Checkout(ctx, input)
}
