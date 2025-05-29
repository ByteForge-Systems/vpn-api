package routes

import (
	"github.com/ByteForge-Systems/vpn-api/internal/transport/handlers"
	"github.com/gin-gonic/gin"
)

func SetupManagementRoutes(router *gin.Engine, handler *handlers.ManagementHandler) {
	management := router.Group("/api/management")
	{
		management.POST("/restart/:node_id", handler.RestartXray)
		management.GET("/status/:node_id", handler.GetXrayStatus)
		management.POST("/start/:node_id", handler.StartXray)
		management.POST("/stop/:node_id", handler.StopXray)

	}
}
