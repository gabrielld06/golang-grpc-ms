package service

import (
	"context"
	"grpc-microsservice/pb/orders"
	"slices"
)

var id = 1
var ordersMem = make([]*orders.Order, 0)

type OrderService struct{}

func NewOrderService() *OrderService {
	return &OrderService{}
}

func (s *OrderService) CreateOrder(ctx context.Context, order *orders.Order) error {
	order.OrderID = int32(id)
	id += 1

	ordersMem = append(ordersMem, order)

	return nil
}

func (s *OrderService) GetOrders(ctx context.Context, customerID int32) ([]*orders.Order, error) {
	if customerID == 0 {
		return ordersMem, nil
	}

	orders := slices.Collect(func(yield func(*orders.Order) bool) {
		for _, o := range ordersMem {
			if o.CustomerID != customerID {
				continue
			}

			if !yield(o) {
				return
			}
		}
	})

	return orders, nil
}
