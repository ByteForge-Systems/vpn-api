package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-api/api/routes"
	"github.com/ByteForge-Systems/vpn-api/utils"
	swaggerFiles "github.com/swaggo/files"
    ginSwagger "github.com/swaggo/gin-swagger"
    _ "github.com/ByteForge-Systems/vpn-api/api/docs"
)

// @title Xray Management API
// @version 1.0
// @description API для управления пользователями и Xray
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.example.com/support
// @contact.email support@example.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /

var NODE_API_BASE_URL string
func init() {
	utils.LoadEnv()
	NODE_API_BASE_URL = utils.GetEnv("NODE_API_BASE_URL")
}

func main() {
	router := gin.Default()

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	routes.SetupUserRoutes(router)
	routes.SetupManagementRoutes(router)

	router.Run(":8080")
}