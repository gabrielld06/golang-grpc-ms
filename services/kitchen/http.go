package main

import (
	"grpc-microsservice/services/kitchen/handlers"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type HttpServer struct {
	addr string
}

func NewHttpServer(addr string) *HttpServer {
	return &HttpServer{
		addr: addr,
	}
}

func NewGRPCClient(addr string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	return conn, err
}

func (s *HttpServer) Run() error {
	router := http.NewServeMux()

	grpcClient, err := NewGRPCClient(":9000")
	if err != nil {
		return err
	}
	defer grpcClient.Close()

	kitchenHandler := handlers.NewKitchenHttpHandler(grpcClient)
	kitchenHandler.RegisterRoutes(router)

	log.Println("Starting http server on", s.addr)

	return http.ListenAndServe(s.addr, router)
}
