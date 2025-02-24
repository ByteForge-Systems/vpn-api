package main

import (
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-api/api/routes"
	"github.com/ByteForge-Systems/vpn-api/utils"
)

var NODE_API_BASE_URL string
func init() {
	utils.LoadEnv()
	NODE_API_BASE_URL = utils.GetEnv("NODE_API_BASE_URL")
}

func main() {
	router := gin.Default()

	routes.SetupUserRoutes(router)
	routes.SetupManagementRoutes(router)

	router.Run(":8080")
}