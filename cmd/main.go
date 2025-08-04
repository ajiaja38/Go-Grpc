package main

import (
	"log"
	"net"

	"go-grpc/cmd/config"
	"go-grpc/cmd/services"
	productPb "go-grpc/pb"

	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func main() {
	netListen, err := net.Listen("tcp", port)

	if err != nil {
		log.Fatalf("Failed to listen: %v", err.Error())
	}

	db := config.ConnectDB()

	grpcServer := grpc.NewServer()
	productService := services.ProductService{DB: db}
	productPb.RegisterProductServiceServer(grpcServer, &productService)

	log.Printf("Starting gRPC server on %s", port)

	if err := grpcServer.Serve(netListen); err != nil {
		log.Fatalf("Failed to serve: %v", err.Error())
	}
}
