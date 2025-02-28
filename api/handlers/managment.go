package handlers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/ByteForge-Systems/vpn-api/client"
	"github.com/ByteForge-Systems/vpn-api/models"
)

// @Summary Перезапустить Xray
// @Description Перезапускает Xray
// @Tags Management
// @Produce json
// @Success 200 {object} models.SuccessResponse
// @Failure 500 {object} models.ErrorResponse
// @Router /api/management/restart [post]
func RestartXray(c *gin.Context) {
	err := client.RestartXray()
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
	status, err := client.GetXrayStatus()
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
    err := client.StartXray()
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
    err := client.StopXray()
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
        return
    }
    c.JSON(http.StatusOK, gin.H{"message": "Xray stopped"})
}