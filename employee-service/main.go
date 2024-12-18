package main

import (
	"hr-system/common/config"
	"hr-system/common/dao/query"
	"hr-system/common/database"
	"hr-system/common/proto"

	"hr-system/employee-service/internal"
	"log"
	"net"

	"google.golang.org/grpc"
)

func main() {
	database.Init(config.Get())
	query.SetDefault(database.GetDB())

	listener, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen on port 50052: %v", err)
	}

	server := grpc.NewServer()
	proto.RegisterEmployeeServiceServer(server, &internal.Server{})

	log.Println("Employee Service running at :50052")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to start Employee Service: %v", err)
	}
}
