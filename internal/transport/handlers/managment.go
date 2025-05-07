package handlers

import (
	"fmt"
	"github.com/ByteForge-Systems/vpn-api/internal/models"
	"net/http"

	_ "github.com/ByteForge-Systems/vpn-api/internal/models"
	"github.com/ByteForge-Systems/vpn-api/internal/node_client"
	"github.com/gin-gonic/gin"
)

// getNodeURLByID получает URL ноды по её ID
func (h *NodeHandler) getNodeURLByID(nodeID string) (string, error) {
	var node models.Node
	err := h.store.DB.Get(&node, "SELECT ip, port FROM nodes WHERE id = $1", nodeID)
	if err != nil {
		return "", fmt.Errorf("failed to find node: %w", err)
	}
	return fmt.Sprintf("http://%s:%d", node.IP, node.Port), nil
}

// @Summary Перезапустить Xray
// @Description Перезапускает Xray на указанной ноде
// @Tags Management
// @Produce json
// @Param node_id path string true "ID ноды"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/management/{node_id}/restart [post]
func (h *NodeHandler) RestartXray(c *gin.Context) {
	nodeID := c.Param("node_id")
	if nodeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "node_id is required"})
		return
	}

	nodeURL, err := h.getNodeURLByID(nodeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = node_client.RestartXray(nodeURL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xray restarted"})
}

// @Summary Получить статус Xray

// @Description Возвращает статус Xray на указанной ноде
// @Tags Management
// @Produce json
// @Param node_id path string true "ID ноды"
// @Success 200 {object} models.StatusResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/management/{node_id}/status [get]
func (h *NodeHandler) GetXrayStatus(c *gin.Context) {
	nodeID := c.Param("node_id")
	if nodeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "node_id is required"})
		return
	}

	nodeURL, err := h.getNodeURLByID(nodeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	status, err := node_client.GetXrayStatus(nodeURL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary Запустить Xray

// @Description Запускает Xray на указанной ноде
// @Tags Management
// @Produce json
// @Param node_id path string true "ID ноды"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/management/{node_id}/start [post]
func (h *NodeHandler) StartXray(c *gin.Context) {
	nodeID := c.Param("node_id")
	if nodeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "node_id is required"})
		return
	}

	nodeURL, err := h.getNodeURLByID(nodeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = node_client.StartXray(nodeURL)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Xray started"})
}

// @Summary Остановить Xray

// @Description Останавливает Xray на указанной ноде
// @Tags Management
// @Produce json
// @Param node_id path string true "ID ноды"
// @Success 200 {object} models.SuccessResponse
// @Failure 400 {object} models.ErrorResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/management/{node_id}/stop [post]
func (h *NodeHandler) StopXray(c *gin.Context) {
	nodeID := c.Param("node_id")
	if nodeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "node_id is required"})
		return
	}

	nodeURL, err := h.getNodeURLByID(nodeID)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = node_client.StopXray(nodeURL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xray stopped"})
}
