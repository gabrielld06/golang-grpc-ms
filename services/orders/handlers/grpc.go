package handlers

import (
	"context"
	"grpc-microsservice/pb/orders"
	"grpc-microsservice/services/orders/types"

	"google.golang.org/grpc"
)

type OrdersGrpcHandler struct {
	ordersService types.OrderService
	orders.UnimplementedOrderServiceServer
}

func NewOrdersGrpcHandler(grpcServer *grpc.Server, orderService types.OrderService) *OrdersGrpcHandler {
	grpcHandler := &OrdersGrpcHandler{
		ordersService: orderService,
	}

	orders.RegisterOrderServiceServer(grpcServer, grpcHandler)

	return grpcHandler
}

func (h *OrdersGrpcHandler) CreateOrder(ctx context.Context, payload *orders.CreateOrderRequest) (*orders.CreateOrderResponse, error) {
	order := &orders.Order{
		OrderID:    0,
		CustomerID: payload.CustomerID,
		ProductID:  payload.ProductID,
		Quantity:   payload.Quantity,
	}

	if err := h.ordersService.CreateOrder(ctx, order); err != nil {
		return nil, err
	}

	return &orders.CreateOrderResponse{
		Status: "success",
	}, nil
}

func (h *OrdersGrpcHandler) GetOrders(ctx context.Context, payload *orders.GetOrdersRequest) (*orders.GetOrdersResponse, error) {
	ordersList, err := h.ordersService.GetOrders(ctx, payload.CustomerID)
	if err != nil {
		return nil, err
	}

	return &orders.GetOrdersResponse{
		Orders: ordersList,
	}, nil
}
