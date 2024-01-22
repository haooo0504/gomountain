package service

import (
	"fmt"
	"gomountain/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddFavoriteMountainRoad
// @Security ApiKeyAuth
// @Summary 最愛的山名及路名
// @Tags 最愛的山名及路名
// @param userId formData string false "用戶ID"
// @param mountainRoadID  formData string false "山名ID"
// @Success 200 {string} json{"code","message"}
// @Router /favorite/addFavoriteMountainRoad [post]
func AddFavoriteMountainRoad(c *gin.Context) {
	userId, err := strconv.ParseUint(c.PostForm("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, // 0 成功 -1失敗
			"message": "不正確的userId"})
		return
	}
	mountainRoadID, err := strconv.ParseUint(c.PostForm("mountainRoadID"), 10, 32)
	if err != nil {

		c.JSON(http.StatusBadRequest, gin.H{"code": -1, // 0 成功 -1失敗
			"message": "不正確的mountainRoadID "})
		return
	}
	AddToFavoritesErr := models.AddToFavorites(uint(userId), uint(mountainRoadID))

	if AddToFavoritesErr != nil {
		fmt.Println(AddToFavoritesErr)
		if AddToFavoritesErr.Error() == "已經加入最愛" {
			c.JSON(http.StatusOK, gin.H{
				"code":    0, // 0 成功 -1失敗
				"message": "已經加入最愛",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, // 0 成功 -1失敗
			"message": "加入最愛失敗",
			"data":    AddToFavoritesErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0, // 0 成功 -1失敗
		"message": "加入最愛成功",
	})
}

// GetFavoriteMountainRoad
// @Security ApiKeyAuth
// @Summary 用戶最愛的山名及路名
// @Tags 最愛的山名及路名
// @param userId query string false "用戶ID"
// @Success 200 {string} json{"code","message"}
// @Router /favorite/getFavoriteMountainRoad [get]
func GetFavoriteMountainRoad(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Query("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, // 0 成功 -1失敗
			"message": "不正確的userId"})
		return
	}

	data, GetFavoritesErr := models.GetUserFavoriteMountainRoads(uint(userId))

	if GetFavoritesErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, // 0 成功 -1失敗
			"message": "加入最愛失敗",
			"data":    GetFavoritesErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0, // 0 成功 -1失敗
		"message": "加入最愛成功",
		"data":    data,
	})
}

// DelFavoriteMountainRoad
// @Security ApiKeyAuth
// @Summary 移除用戶最愛的山名及路名
// @Tags 最愛的山名及路名
// @param userId query string false "用戶ID"
// @param mountainRoadID query string false "山名ID"
// @Success 200 {string} json{"code","message"}
// @Router /favorite/delFavoriteMountainRoad [delete]
func DelFavoriteMountainRoad(c *gin.Context) {
	userId, err := strconv.ParseUint(c.Query("userId"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, // 0 成功 -1失敗
			"message": "不正確的userId"})
		return
	}

	mountainRoadID, err := strconv.ParseUint(c.Query("mountainRoadID"), 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, // 0 成功 -1失敗
			"message": "不正確的mountainRoadID"})
		return
	}

	DelFavoritesErr := models.RemoveFromFavorites(uint(userId), uint(mountainRoadID))

	if DelFavoritesErr != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1, // 0 成功 -1失敗
			"message": "移除最愛失敗",
			"data":    DelFavoritesErr.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    0, // 0 成功 -1失敗
		"message": "移除最愛成功",
	})
}
