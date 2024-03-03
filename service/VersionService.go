package service

import (
	"gomountain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetVersion
// @Summary 獲取版本號
// @Tags 獲取版本號
// @Success 200 {string} json{"code","message"}
// @Router /version/getVersion [get]
func GetVersion(c *gin.Context) {
	data, err := models.GetVersion()

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Version retrieved successfully", "data": data})
}
