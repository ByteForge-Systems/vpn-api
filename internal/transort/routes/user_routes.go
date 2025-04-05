package routes

import (
	"github.com/ByteForge-Systems/vpn-api/internal/transort/handlers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine) {
	user := router.Group("/api/user")
	{
		user.POST("/", handlers.AddUser)
		user.DELETE("/:id", handlers.RemoveUser)
		user.GET("/:id/link", handlers.GenerateVLESSLink)
		user.GET("/all", handlers.GetAllUsers)
	}
}
