package handlers

import (
	"grpc-microsservice/pb/orders"
	"html/template"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"google.golang.org/grpc"
)

type KitchenHttpHandler struct {
	grpcClientConn *grpc.ClientConn
}

func NewKitchenHttpHandler(grpcClientConn *grpc.ClientConn) *KitchenHttpHandler {
	grpcHandler := &KitchenHttpHandler{
		grpcClientConn: grpcClientConn,
	}

	return grpcHandler
}

func (h *KitchenHttpHandler) RegisterRoutes(router *http.ServeMux) {
	router.HandleFunc("GET /orders", h.Orders)
	router.HandleFunc("GET /customer/{id}", h.CustomerOrders)
}

func (h *KitchenHttpHandler) Orders(w http.ResponseWriter, r *http.Request) {
	c := orders.NewOrderServiceClient(h.grpcClientConn)

	c.CreateOrder(r.Context(), &orders.CreateOrderRequest{
		CustomerID: rand.Int31n(256),
		ProductID:  rand.Int31n(256),
		Quantity:   rand.Int31n(8),
	})

	ordersList, err := c.GetOrders(r.Context(), &orders.GetOrdersRequest{
		CustomerID: 0,
	})
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.New("orders").Parse(ordersTemplate))

	if err := t.Execute(w, ordersList.GetOrders()); err != nil {
		log.Fatal(err)
	}
}

func (h *KitchenHttpHandler) CustomerOrders(w http.ResponseWriter, r *http.Request) {
	c := orders.NewOrderServiceClient(h.grpcClientConn)

	customerIdStr := r.PathValue("id")
	var createCustomerId int32
	var getCustomerId int32
	if cId, err := strconv.ParseInt(customerIdStr, 10, 32); err != nil {
		// id not provided
		createCustomerId = rand.Int31n(256)
		getCustomerId = 0
	} else {
		createCustomerId = int32(cId)
		getCustomerId = int32(cId)
	}

	c.CreateOrder(r.Context(), &orders.CreateOrderRequest{
		CustomerID: createCustomerId,
		ProductID:  rand.Int31n(256),
		Quantity:   rand.Int31n(8),
	})

	ordersList, err := c.GetOrders(r.Context(), &orders.GetOrdersRequest{
		CustomerID: getCustomerId,
	})
	if err != nil {
		log.Fatal(err)
	}

	t := template.Must(template.New("orders").Parse(ordersTemplate))

	if err := t.Execute(w, ordersList.GetOrders()); err != nil {
		log.Fatal(err)
	}
}

var ordersTemplate = `
<!DOCTYPE html>
<html>
<head>
    <title>Kitchen Orders</title>
</head>
<body>
    <h1>Orders List</h1>
    <table border="1">
        <tr>
            <th>Order ID</th>
            <th>Customer ID</th>
            <th>Quantity</th>
        </tr>
        {{range .}}
        <tr>
            <td>{{.OrderID}}</td>
            <td>{{.CustomerID}}</td>
            <td>{{.Quantity}}</td>
        </tr>
        {{end}}
    </table>
</body>
</html>`
