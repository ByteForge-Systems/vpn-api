package handlers

import (
	_ "github.com/ByteForge-Systems/vpn-api/internal/models"
	"github.com/ByteForge-Systems/vpn-api/internal/node_client"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
)

// @Summary Добавить пользователя
// @Description Добавляет нового пользователя в систему
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.User true "Данные пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /api/user/ [post]

func AddUser(c *gin.Context) {
	newUUID := uuid.New().String()

	_, err := node_client.AddUser(newUUID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"uuid": newUUID})
}

// @Summary Удалить пользователя
// @Description Удаляет пользователя по ID
// @Tags User
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/user/{id} [delete]

func RemoveUser(c *gin.Context) {
	userID := c.Param("id")
	err := node_client.RemoveUser(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// @Summary Получить всех пользователей
// @Description Возвращает список всех пользователей
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Router /api/user/all [get]

func GetAllUsers(c *gin.Context) {
	clients, err := node_client.GetAllUsers() // слайс
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": clients})
}

// @Summary Сгенерировать VLESS-ссылку
// @Description Генерирует VLESS-ссылку для пользователя по ID
// @Tags User
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 200 {object} models.VLESSLink
// @Failure 404 {object} models.ErrorResponse
// @Router /api/user/{id}/link [get]

func GenerateVLESSLink(c *gin.Context) {
	userID := c.Param("id")
	link, err := node_client.GenerateVLESSLink(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"link": link})
}
