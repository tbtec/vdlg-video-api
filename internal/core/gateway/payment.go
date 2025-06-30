package gateway

import (
	"context"
	"log/slog"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
	"github.com/tbtec/tremligeiro/internal/infra/external"
)

const (
	SPONSOR_ID = 96469944
)

type PaymentGateway struct {
	paymentService    external.IPaymentService
	paymentRepository repository.IPaymentRepository
}

func NewPaymentGateway(paymentService external.IPaymentService,
	paymentRepository repository.IPaymentRepository) *PaymentGateway {
	return &PaymentGateway{
		paymentService:    paymentService,
		paymentRepository: paymentRepository,
	}
}

func (gtw *PaymentGateway) RequestPayment(ctx context.Context, order entity.Order, orderProduct []entity.OrderProduct, metadata dto.MetaData) error {

	items := make([]dto.PaymentItemCheckoutProduct, 0)

	for _, op := range orderProduct {
		items = append(items, dto.PaymentItemCheckoutProduct{
			ProductId: op.ProductID,
			Quantity:  int(op.Quantity),
		})
	}

	paymentRequest := dto.PaymentCheckout{
		OrderId:     order.ID,
		TotalAmount: order.TotalAmount,
		Products:    items,
	}

	slog.InfoContext(ctx, "Requesting payment...")

	response, err := gtw.paymentService.RequestPayment(ctx, paymentRequest)
	if err != nil {
		slog.ErrorContext(ctx, "❌ Error requesting payment", "error", err)
		return err
	}

	slog.InfoContext(ctx, "✅ Request payment succesfully", "response", response)

	return nil
}
