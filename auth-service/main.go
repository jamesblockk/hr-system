package main

import (
	"hr-system/auth-service/internal"
	"hr-system/common/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {

	listener, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterAuthServiceServer(server, &internal.Server{})

	log.Println("Auth Service running at :50051")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start Auth Service: %v", err)
	}
}
