package routes

import (
	"github.com/ByteForge-Systems/vpn-api/internal/transport/handlers"
	"github.com/gin-gonic/gin"
)

func SetupManagementRoutes(router *gin.Engine, handler *handlers.NodeHandler) {
	management := router.Group("/api/management")
	{
		management.POST("/restart", handler.RestartXray)
		management.GET("/status", handler.GetXrayStatus)
		management.POST("/start", handler.StartXray)
		management.POST("/stop", handler.StopXray)

	}
}
