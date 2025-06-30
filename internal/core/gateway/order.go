package gateway

import (
	"context"
	"time"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/infra/database/model"
	"github.com/tbtec/tremligeiro/internal/infra/database/repository"
)

type OrderGateway struct {
	orderRepository repository.IOrderRepository
}

func NewOrderGateway(orderRepository repository.IOrderRepository) *OrderGateway {
	return &OrderGateway{
		orderRepository: orderRepository,
	}
}

func (gtw *OrderGateway) Create(ctx context.Context, order *entity.Order) error {

	orderModel := model.Order{
		ID:          order.ID,
		CustomerId:  order.CustomerId,
		Status:      string(order.Status),
		TotalAmount: order.TotalAmount,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
	}

	err := gtw.orderRepository.Create(ctx, &orderModel)

	if err != nil {
		return err
	}

	return nil
}

func (gtw *OrderGateway) Update(ctx context.Context, order *entity.Order) error {

	orderModel := model.Order{
		ID:          order.ID,
		CustomerId:  order.CustomerId,
		Status:      string(order.Status),
		TotalAmount: order.TotalAmount,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   time.Now().UTC(),
	}

	err := gtw.orderRepository.Update(ctx, &orderModel)

	if err != nil {
		return err
	}

	return nil
}

func (gtw *OrderGateway) FindOne(ctx context.Context, id string) (*entity.Order, error) {
	orderModel, err := gtw.orderRepository.FindOne(ctx, id)
	if err != nil {
		return nil, err
	}

	order := entity.Order{
		ID:          orderModel.ID,
		CustomerId:  orderModel.CustomerId,
		Status:      entity.OrderStatus(orderModel.Status),
		TotalAmount: orderModel.TotalAmount,
		CreatedAt:   orderModel.CreatedAt,
		UpdatedAt:   orderModel.UpdatedAt,
	}

	return &order, nil
}

func (gtw *OrderGateway) Find(ctx context.Context) ([]entity.Order, error) {
	orderModels, err := gtw.orderRepository.Find(ctx)
	if err != nil {
		return nil, err
	}

	orders := []entity.Order{}
	for _, orderModel := range orderModels {
		order := entity.Order{
			ID:          orderModel.ID,
			CustomerId:  orderModel.CustomerId,
			Status:      entity.OrderStatus(orderModel.Status),
			TotalAmount: orderModel.TotalAmount,
			CreatedAt:   orderModel.CreatedAt,
			UpdatedAt:   orderModel.UpdatedAt,
		}
		orders = append(orders, order)
	}

	return orders, nil
}

// func (gtw *OrderGateway) PublishMessage(ctx context.Context, order *entity.Order) error {
// 	orderDto := dto.Order{
// 		ID:          order.ID,
// 		CustomerId:  order.CustomerId,
// 		Status:      string(order.Status),
// 		TotalAmount: order.TotalAmount,
// 		CreatedAt:   order.CreatedAt,
// 		UpdatedAt:   order.UpdatedAt,
// 	}

// 	err := gtw.producerService.PublishMessage(ctx, orderDto)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
