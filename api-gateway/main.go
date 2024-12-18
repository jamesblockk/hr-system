package main

import (
	"hr-system/api-gateway/client"
	"hr-system/api-gateway/routes"
	"hr-system/api-gateway/service"

	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 初始化 gRPC 客戶端
	authClient, err := client.NewAuthServiceClient("auth-service:50051")
	if err != nil {
		panic(err)
	}
	defer authClient.Close()

	employeeClient, err := client.NewEmployeeServiceClient("employee-service:50052")
	if err != nil {
		panic(err)
	}
	defer employeeClient.Close()

	s := service.New(
		service.NewAuthService(authClient),
		service.NewEmployeeService(employeeClient),
	)

	routes.RegisterRoutes(r, s)

	err = r.Run(":8081")
	if err != nil {
		log.Fatal("Unable to start the server: ", err)
	}
}
