package handlers

import (
	"net/http"
	"strconv"

	"github.com/ByteForge-Systems/vpn-api/internal/db/nodes"
	"github.com/ByteForge-Systems/vpn-api/internal/models"
	"github.com/gin-gonic/gin"
)

type NodeHandler struct {
	store *nodes.NodeStore
}

func NewNodeHandler(store *nodes.NodeStore) *NodeHandler {
	return &NodeHandler{store: store}
}

// CreateNode godoc
// @Summary Создать новый узел
// @Description Создает новый VPN-узел с указанными данными
// @Tags узлы
// @Accept json
// @Produce json
// @Param node body models.Node true "Данные узла"
// @Success 201 {object} models.Node "Узел создан, пароль не возвращается"
// @Failure 400 {object} map[string]string "Неверный формат запроса"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/nodes [post]
func (h *NodeHandler) CreateNode(c *gin.Context) {
	var node models.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}

	if err := h.store.Create(&node); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, node)
}

// GetNode godoc
// @Summary Получить узел по ID
// @Description Получает VPN-узел по его идентификатору
// @Tags узлы
// @Produce json
// @Param id path int true "ID узла"
// @Success 200 {object} models.Node "Детали узла, пароль не возвращается"
// @Failure 400 {object} map[string]string "Неверный ID узла"
// @Failure 404 {object} map[string]string "Узел не найден"
// @Router /api/nodes/{id} [get]
func (h *NodeHandler) GetNode(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid node ID"})
		return
	}

	node, err := h.store.Get(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, node)
}

// ListNodes godoc
// @Summary Список всех узлов
// @Description Получает список всех VPN-узлов
// @Tags узлы
// @Produce json
// @Success 200 {array} models.Node "Список узлов, пароли не возвращаются"
// @Failure 500 {object} map[string]string "Внутренняя ошибка сервера"
// @Router /api/nodes [get]
func (h *NodeHandler) ListNodes(c *gin.Context) {
	nodes, err := h.store.List()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, nodes)
}

// UpdateNode godoc
// @Summary Обновить узел
// @Description Обновляет существующий VPN-узел по его ID
// @Tags узлы
// @Accept json
// @Produce json
// @Param id path int true "ID узла"
// @Param node body models.Node true "Обновленные данные узла"
// @Success 200 {object} models.Node "Обновленный узел, пароль не возвращается"
// @Failure 400 {object} map[string]string "Неверный формат запроса или ID узла"
// @Failure 404 {object} map[string]string "Узел не найден"
// @Router /api/nodes/{id} [put]
func (h *NodeHandler) UpdateNode(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid node ID"})
		return
	}

	var node models.Node
	if err := c.ShouldBindJSON(&node); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request payload"})
		return
	}
	node.ID = id

	if err := h.store.Update(&node); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, node)
}

// DeleteNode godoc
// @Summary Удалить узел
// @Description Удаляет VPN-узел по его ID
// @Tags узлы
// @Param id path int true "ID узла"
// @Success 204 "Нет содержимого"
// @Failure 400 {object} map[string]string "Неверный ID узла"
// @Failure 404 {object} map[string]string "Узел не найден"
// @Router /api/nodes/{id} [delete]
func (h *NodeHandler) DeleteNode(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid node ID"})
		return
	}

	if err := h.store.Delete(id); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
