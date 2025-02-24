package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-api/client"
	"github.com/google/uuid"
)


func AddUser(c *gin.Context) {
    newUUID := uuid.New().String()
    
    _, err := client.AddUser(newUUID)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"uuid": newUUID})
}

func RemoveUser(c *gin.Context) {
	userID := c.Param("id")
	err := client.RemoveUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

func GetAllUsers(c *gin.Context) {
    clients, err := client.GetAllUsers() // слайс
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }

    c.JSON(http.StatusOK, gin.H{"users": clients})
}

func GenerateVLESSLink(c *gin.Context) {
	userID := c.Param("id")
	link, err := client.GenerateVLESSLink(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"link": link})
}