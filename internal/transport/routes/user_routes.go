package routes

import (
	_ "github.com/ByteForge-Systems/vpn-api/internal/node_client"
	"github.com/ByteForge-Systems/vpn-api/internal/transport/handlers"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/jmoiron/sqlx"
)

func SetupUserRoutes(router *gin.Engine, db *sqlx.DB) {
	userHandler := handlers.NewUserHandler(db)
	user := router.Group("/api/users")
	{
		user.POST("/", userHandler.AddUser)
		user.DELETE("/:id", userHandler.RemoveUser)
		user.GET("/:id/link", userHandler.GenerateVLESSLink)
		user.GET("/all", userHandler.GetAllUsers)
	}
}
