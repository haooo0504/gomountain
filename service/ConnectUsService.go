package service

import (
	"gomountain/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// ConnectToUs
// @Summary 聯絡我們
// @Tags 聯絡我們
// @param userId formData string false "用戶ID"
// @param userName formData string false "用戶姓名"
// @param gmail formData string false "gmail"
// @param userComment formData string false "留言內容"
// @Success 200 {string} json{"code","message"}
// @Router /connect/connectToUs [post]
func ConnectToUsHandler(c *gin.Context) {
	userId := c.PostForm("userId")
	userName := c.PostForm("userName")
	gmail := c.PostForm("gmail")
	userComment := c.PostForm("userComment")

	res, err := models.ConnectToUs(userId, userName, gmail, userComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to add comment"})
		return
	}

	// 返回成功響應
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Comment added successfully", "data": res})
}
