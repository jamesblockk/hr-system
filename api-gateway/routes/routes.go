package routes

import (
	"hr-system/api-gateway/middlewares"
	"hr-system/api-gateway/service"

	swaggerFiles "github.com/swaggo/files" // swagger embed files

	"hr-system/api-gateway/handler"
	_ "hr-system/docs" // 這裡引入您本地的 docs package

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server Petstore server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8081
// @BasePath /api/
func RegisterRoutes(router *gin.Engine, s *service.Service) {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	api := router.Group("/api")
	{

		api.GET("/health", handler.Health)

		api.POST("/register", handler.Register(s))

		api.POST("/login", handler.Login(s))

		authorizedGroup := api.Group("", middlewares.JWTAuth())
		{
			authorizedGroup.GET("/employees/:id", handler.GetEmployeeByID(s))
		}
	}
}
