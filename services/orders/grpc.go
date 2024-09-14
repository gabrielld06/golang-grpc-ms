package main

import (
	"grpc-microsservice/services/orders/handlers"
	"grpc-microsservice/services/orders/service"
	"log"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *GRPCServer {
	return &GRPCServer{
		addr: addr,
	}
}

func (s *GRPCServer) Run() error {
	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		return err
	}

	grpcServer := grpc.NewServer()

	// register grpc services
	orderService := service.NewOrderService()
	handlers.NewOrdersGrpcHandler(grpcServer, orderService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(listener)
}
