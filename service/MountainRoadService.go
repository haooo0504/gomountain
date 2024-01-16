package service

import (
	"gomountain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetMountainRoad
// @Security ApiKeyAuth
// @Summary 取得山名路名
// @Tags 取得山名路名
// @Success 200 {string} json{"code","message"}
// @Router /mountainRoad/getMountainRoad [get]
func GetMountainRoad(c *gin.Context) {
	data, err := models.GetMountainRoad()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1, // 0 成功 -1失敗
			"message": "取得山名路名失敗",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0, // 0 成功 -1失敗
		"message": "取得山名路名成功",
		"data":    data,
	})

}
