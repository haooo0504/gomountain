package service

import (
	"gomountain/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// AddComment
// @Security ApiKeyAuth
// @Summary 貼文留言
// @Tags 貼文留言
// @param userId formData string false "用戶ID"
// @param postId formData string false "貼文ID"
// @param userComment formData string false "留言內容"
// @Success 200 {string} json{"code","message"}
// @Router /comment/addComment [post]
func AddCommentHandler(c *gin.Context) {
	// 从表单中获取数据
	userIdStr := c.PostForm("userId")
	postIdStr := c.PostForm("postId")
	userComment := c.PostForm("userComment")

	// 将字符串转换为 uint
	userId, err := strconv.ParseUint(userIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid user ID"})
		return
	}

	postId, err := strconv.ParseUint(postIdStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid post ID"})
		return
	}

	// 调用添加留言的服务
	_, err = models.AddComment(uint(userId), uint(postId), userComment)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to add comment"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Comment added successfully"})
}

// GetComment
// @Security ApiKeyAuth
// @Summary 獲取貼文留言
// @Tags 貼文留言
// @param postId query string false "貼文ID"
// @Success 200 {string} json{"code","message"}
// @Router /comment/getComment [get]
func GetComment(c *gin.Context) {
	postId := c.Query("postId")

	curPostId, err := strconv.ParseUint(postId, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": 400, "message": "Invalid post ID"})
		return
	}

	// 调用添加留言的服务
	data, err := models.GetCommentsByPostID(uint(curPostId))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": 500, "message": "Failed to add comment"})
		return
	}

	// 返回成功响应
	c.JSON(http.StatusOK, gin.H{"code": 200, "message": "Comment added successfully", "data": data})
}
