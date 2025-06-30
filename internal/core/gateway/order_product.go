package gateway

import (
	"context"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
)

type OrderProductGateway struct {
	orderProductRepository repository.IOrderProductRepository
}

func NewOrderProductGateway(orderProductRepository repository.IOrderProductRepository) *OrderProductGateway {
	return &OrderProductGateway{
		orderProductRepository: orderProductRepository,
	}
}

func (gtw *OrderProductGateway) Create(ctx context.Context, orderProduct entity.OrderProduct) error {

	orderProductModel := model.OrderProduct{
		ID:          orderProduct.ID,
		OrderID:     orderProduct.OrderID,
		ProductID:   orderProduct.ProductID,
		Quantity:    orderProduct.Quantity,
		Amount:      orderProduct.Amount,
		TotalAmount: orderProduct.TotalAmount,
		CreatedAt:   orderProduct.CreatedAt,
	}

	err := gtw.orderProductRepository.Create(ctx, &orderProductModel)

	if err != nil {
		return err
	}

	return nil

}

func (gtw *OrderProductGateway) FindByOrderId(ctx context.Context, orderId string) ([]entity.OrderProduct, error) {

	orderProductsModel, err := gtw.orderProductRepository.FindByOrderId(ctx, orderId)
	if err != nil {
		return nil, err
	}

	orderProduct := []entity.OrderProduct{}
	for _, opModel := range orderProductsModel {
		orderProduct = append(orderProduct, entity.OrderProduct{
			ID:          opModel.ID,
			OrderID:     opModel.OrderID,
			ProductID:   opModel.ProductID,
			Quantity:    opModel.Quantity,
			Amount:      opModel.Amount,
			TotalAmount: opModel.TotalAmount,
		})
	}

	return orderProduct, nil
}
