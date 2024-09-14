package main

import "log"

func main() {
	httpServer := NewHttpServer(":8000")
	go func() {
		log.Fatal(httpServer.Run())
	}()

	grpcServer := NewGRPCServer(":9000")
	log.Fatal(grpcServer.Run())
}
