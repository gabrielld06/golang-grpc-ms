package main

import (
	"log"
)

func main() {
	httpServer := NewHttpServer(":3000")
	log.Fatal(httpServer.Run())
}
