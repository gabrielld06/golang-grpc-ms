package handlers

import (
	"grpc-microsservice/pb/orders"
	"grpc-microsservice/services/orders/types"
	"grpc-microsservice/services/util"
	"net/http"
	"strconv"
)

type OrdersHttpHandler struct {
	ordersService types.OrderService
}

func NewOrdersHttpHandler(orderService types.OrderService) *OrdersHttpHandler {
	grpcHandler := &OrdersHttpHandler{
		ordersService: orderService,
	}

	return grpcHandler
}

func (h *OrdersHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("POST /orders", h.CreateOrder)
	router.HandleFunc("GET /orders", h.GetOrders)
}

func (h *OrdersHttpHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var payload orders.CreateOrderRequest
	if err := util.ParseJSON(r, &payload); err != nil {
		util.WriteError(w, 422, err)
		return
	}

	order := &orders.Order{
		OrderID:    1,
		CustomerID: payload.CustomerID,
		ProductID:  payload.ProductID,
		Quantity:   payload.Quantity,
	}

	if err := h.ordersService.CreateOrder(r.Context(), order); err != nil {
		util.WriteError(w, 500, err)
		return
	}

	util.WriteJSON(w, 201, &orders.CreateOrderResponse{
		Status: "success",
	})
}

func (h *OrdersHttpHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	customerIdStr := r.URL.Query().Get("customerId")
	var customerId int32
	if cId, err := strconv.ParseInt(customerIdStr, 10, 32); err != nil {
		customerId = 0
	} else {
		customerId = int32(cId)
	}

	ordersList, err := h.ordersService.GetOrders(r.Context(), customerId)
	if err != nil {
		return
	}

	util.WriteJSON(w, 200, &orders.GetOrdersResponse{
		Orders: ordersList,
	})
}
