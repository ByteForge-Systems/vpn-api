package handlers

import (
	_ "github.com/ByteForge-Systems/vpn-api/internal/models"
	"github.com/ByteForge-Systems/vpn-api/internal/node_client"
	"github.com/gin-gonic/gin"
	"net/http"
)

// @Summary Перезапустить Xray
// @Description Перезапускает Xray
// @Tags Management
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/management/restart [post]

func RestartXray(c *gin.Context) {
	err := node_client.RestartXray()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Xray restarted"})
}

// @Summary Получить статус Xray
// @Description Возвращает текущий статус Xray
// @Tags Management
// @Produce json
// @Success 200 {object} models.XrayStatus
// @Router /api/management/status [get]

func GetXrayStatus(c *gin.Context) {
	status, err := node_client.GetXrayStatus()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": status})
}

// @Summary Запустить Xray
// @Description Запускает Xray
// @Tags Management
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/management/start [post]

func StartXray(c *gin.Context) {
	err := node_client.StartXray()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Xray started"})
}

// @Summary Остановить Xray
// @Description Останавливает Xray
// @Tags Management
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/management/stop [post]

func StopXray(c *gin.Context) {
	err := node_client.StopXray()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Xray stopped"})
}
