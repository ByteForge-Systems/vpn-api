package routes

import (
	"github.com/ByteForge-Systems/vpn-api/internal/transort/handlers"
	"github.com/gin-gonic/gin"
)

func SetupManagementRoutes(router *gin.Engine) {
	management := router.Group("/api/management")
	{
		management.POST("/restart", handlers.RestartXray)
		management.GET("/status", handlers.GetXrayStatus)
		management.POST("/start", handlers.StartXray)
		management.POST("/stop", handlers.StopXray)
	}
}
