package main

import (
	"fmt"
	"github.com/ByteForge-Systems/vpn-api/internal/db"
	"log"
	_ "github.com/ByteForge-Systems/vpn-api/api/docs"
	"github.com/ByteForge-Systems/vpn-api/internal/config"
	"github.com/ByteForge-Systems/vpn-api/internal/db/nodes"
	"github.com/ByteForge-Systems/vpn-api/internal/transport/handlers"
	"github.com/ByteForge-Systems/vpn-api/internal/transport/routes"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
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

func main() {
	// Загружаем конфигурацию
	cfg := config.LoadConfig()

	// Подключаемся к базе данных
	dbConn, err := db.Connect(cfg)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer dbConn.Close()

	// Настраиваем Gin
	router := gin.Default()

	// Swagger UI
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Инициализируем хранилище и хендлер для узлов
	nodeStore := nodes.NewNodeStore(dbConn)
	nodesHandler := handlers.NewNodeHandler(nodeStore)

	// Настраиваем маршруты
	routes.SetupUserRoutes(router, dbConn)
	routes.RegisterNodesRoutes(router, nodesHandler)

	// Запускаем сервер с портом из конфига
	addr := fmt.Sprintf(":%s", cfg.Port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
