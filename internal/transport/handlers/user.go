package handlers

import (
	"database/sql"
	"fmt"
	"regexp"

	"github.com/ByteForge-Systems/vpn-api/internal/models"
	"github.com/ByteForge-Systems/vpn-api/internal/node_client"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"net/http"
)

type UserHandler struct {
	db *sqlx.DB
}

func NewUserHandler(db *sqlx.DB) *UserHandler {
	return &UserHandler{db: db}
}

// @Summary Добавить пользователя
// @Description Добавляет нового пользователя в систему
// @Tags User
// @Accept json
// @Produce json
// @Param user body models.UserRequest true "Данные пользователя"
// @Success 200 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router jupy

func (h *UserHandler) AddUser(c *gin.Context) {
	var req models.UserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}

	// Валидация MAC-адреса
	if !isValidMacAddress(req.MacAddress) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid MAC address"})
		return
	}

	// Проверка, существует ли пользователь с таким MAC-адресом
	var exists bool
	err := h.db.Get(&exists, "SELECT EXISTS(SELECT 1 FROM users WHERE mac_address = $1)", req.MacAddress)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to check existing user"})
		return
	}
	if exists {
		c.JSON(http.StatusBadRequest, gin.H{"error": "user with this MAC address already exists"})
		return
	}

	// Пока что выбираем фиксированную ноду с id=1
	var node models.Node
	err = h.db.Get(&node, `
        SELECT id, ip, port, is_online
        FROM nodes
        WHERE id = 1
    `)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "node with id=1 not found"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to select node: %v", err)})
		return
	}

	// Проверяем, что нода онлайн
	if !node.IsOnline {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "selected node is not online"})
		return
	}

	// Генерация UUID
	newUUID := uuid.New()

	// Создание пользователя в базе
	user := &models.User{
		UUID:       newUUID,
		MacAddress: req.MacAddress,
		NodeID:     int(node.ID),
		Status:     "active",
	}

	// Начинаем транзакцию
	tx, err := h.db.Beginx()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to start transaction"})
		return
	}
	defer tx.Rollback()

	err = tx.QueryRowx(`
        INSERT INTO users (uuid, mac_address, node_id, status, created_at, updated_at)
        VALUES ($1, $2, $3, $4, NOW(), NOW())
        RETURNING id, created_at, updated_at
    `, user.UUID, user.MacAddress, user.NodeID, user.Status).StructScan(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create user"})
		return
	}

	// Формируем URL ноды
	nodeURL := fmt.Sprintf("http://%s:%d", node.IP, node.Port)

	// Отправляем запрос на ноду
	_, err = node_client.AddUser(nodeURL, newUUID.String())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to add user to node: %v", err)})
		return
	}

	// Коммитим транзакцию
	if err := tx.Commit(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to commit transaction"})
		return
	}

	c.JSON(http.StatusOK, user)
}

// @Summary Удалить пользователя
// @Description Удаляет пользователя по ID
// @Tags User
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 200 {object} models.SuccessResponse
// @Failure 404 {object} models.ErrorResponse
// @Router /api/users/{id} [delete]
func (h *UserHandler) RemoveUser(c *gin.Context) {
	userID := c.Param("id")

	// Ищем пользователя
	var user models.User
	err := h.db.Get(&user, "SELECT * FROM users WHERE uuid = $1", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find user"})
		return
	}

	// Ищем ноду
	var node models.Node
	err = h.db.Get(&node, "SELECT ip, port FROM nodes WHERE id = $1", user.NodeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find node"})
		return
	}

	// Формируем URL ноды
	nodeURL := fmt.Sprintf("http://%s:%d", node.IP, node.Port)

	// Удаляем пользователя с ноды
	err = node_client.RemoveUser(nodeURL, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to remove user from node: %v", err)})
		return
	}

	// Удаляем пользователя из базы
	_, err = h.db.Exec("DELETE FROM users WHERE uuid = $1", userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted"})
}

// @Summary Получить всех пользователей
// @Description Возвращает список всех пользователей
// @Tags User
// @Produce json
// @Success 200 {array} models.User
// @Router /api/users/all [get]
func (h *UserHandler) GetAllUsers(c *gin.Context) {
	var users []models.User
	err := h.db.Select(&users, "SELECT * FROM users")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch users"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"users": users})
}

// @Summary Сгенерировать VLESS-ссылку
// @Description Генерирует VLESS-ссылку для пользователя по ID
// @Tags User
// @Produce json
// @Param id path string true "ID пользователя"
// @Success 200 {object} models.VLESSLink
// @Failure 404 {object} models.ErrorResponse
// @Router /api/users/{id}/link [get]
func (h *UserHandler) GenerateVLESSLink(c *gin.Context) {
	userID := c.Param("id")

	// Ищем пользователя
	var user models.User
	err := h.db.Get(&user, "SELECT * FROM users WHERE uuid = $1", userID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find user"})
		return
	}

	// Ищем ноду
	var node models.Node
	err = h.db.Get(&node, "SELECT ip, port FROM nodes WHERE id = $1", user.NodeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to find node"})
		return
	}

	// Формируем URL ноды
	nodeURL := fmt.Sprintf("http://%s:%d", node.IP, node.Port)

	// Генерируем ссылку
	link, err := node_client.GenerateVLESSLink(nodeURL, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("failed to generate VLESS link: %v", err)})
		return
	}

	c.JSON(http.StatusOK, gin.H{"link": link})
}

func isValidMacAddress(mac string) bool {
	matched, _ := regexp.MatchString(`^([0-9A-Fa-f]{2}[:-]){5}([0-9A-Fa-f]{2})$`, mac)
	return matched
}
