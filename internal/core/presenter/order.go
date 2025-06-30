package presenter

import (
	"sort"

	"github.com/tbtec/tremligeiro/internal/core/domain/entity"
	"github.com/tbtec/tremligeiro/internal/dto"
)

type OrderPresenter struct {
}

func NewOrderPresenter() *OrderPresenter {
	return &OrderPresenter{}
}

func (presenter *OrderPresenter) BuildOrderCreateResponse(order entity.Order, paymentId *string) dto.Order {
	return dto.Order{
		ID:          order.ID,
		CustomerId:  order.CustomerId,
		Status:      string(order.Status),
		TotalAmount: order.TotalAmount,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		MetaData: dto.MetaData{
			PaymentId: paymentId,
		},
	}
}

func (presenter *OrderPresenter) BuildOrderContentResponse(orders []entity.Order) dto.OrderContent {
	response := []dto.Order{}

	for i := len(orders) - 1; i >= 0; i-- {
		order := orders[i]
		if order.Status == "FINALIZED" {
			orders = append(orders[:i],
				orders[i+1:]...)
		}
	}

	sort.Slice(orders, func(i, j int) bool {
		return orders[i].CreatedAt.Before(orders[j].CreatedAt)
	})

	for _, order := range orders {
		response = append(response, presenter.BuildOrderCreateResponse(order, nil))
	}

	return dto.OrderContent{Content: response}
}

func (presenter *OrderPresenter) BuildOrderDetailsCreateResponse(order entity.Order,
	orderProducts []entity.OrderProduct) dto.OrderDetails {

	orderProductsResponse := []dto.OrderProduct{}

	for _, op := range orderProducts {
		orderProductsResponse = append(orderProductsResponse, dto.OrderProduct{
			ID:        op.ID,
			ProductID: op.ProductID,
			Quantity:  op.Quantity,
		})
	}

	return dto.OrderDetails{
		ID:            order.ID,
		CustomerId:    order.CustomerId,
		Status:        string(order.Status),
		TotalAmount:   order.TotalAmount,
		CreatedAt:     order.CreatedAt,
		UpdatedAt:     order.UpdatedAt,
		OrderProducts: orderProductsResponse,
	}
}
