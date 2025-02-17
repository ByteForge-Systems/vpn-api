package main

import (
	"github.com/gin-gonic/gin"
	// "github.com/ByteForge-Systems/vpn-api/api/routes"
	// "github.com/ByteForge-Systems/vpn-api/utils"
)

func main() {
	// utils.LoadEnv()

	router := gin.Default()

	// routes.SetupUserRoutes(router)
	// routes.SetupManagementRoutes(router)

	router.Run(":8080")
}