package main

import (
	"grpc-microsservice/services/orders/handlers"
	"grpc-microsservice/services/orders/service"
	"log"
	"net/http"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

func (s *HttpServer) Run() error {
	router := http.NewServeMux()

	orderService := service.NewOrderService()
	orderHandler := handlers.NewOrdersHttpHandler(orderService)
	orderHandler.RegisterRoutes(router)

	log.Println("Starting http server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
