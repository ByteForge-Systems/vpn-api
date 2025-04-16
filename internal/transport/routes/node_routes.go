package routes

import (
	"github.com/ByteForge-Systems/vpn-api/internal/transport/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterNodesRoutes(r *gin.Engine, handler *handlers.NodeHandler) {
	nodes := r.Group("api/nodes")
	{
		nodes.POST("", handler.CreateNode)
		nodes.GET("", handler.ListNodes)
		nodes.GET("/:id", handler.GetNode)
		nodes.PUT("/:id", handler.UpdateNode)
		nodes.DELETE("/:id", handler.DeleteNode)
	}
}
