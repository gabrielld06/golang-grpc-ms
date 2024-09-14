package types

import (
	"context"
	"grpc-microsservice/pb/orders"
)

type OrderService interface {
	CreateOrder(context.Context, *orders.Order) error
	GetOrders(context.Context, int32) ([]*orders.Order, error)
}
