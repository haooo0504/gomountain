package models

import (
	"fmt"
	"gomountain/utils"
	"time"

	"gorm.io/gorm"
)

type PostInfo struct {
	gorm.Model
	Author    string `gorm:"type:varchar(100)"`
	AuthorId  string
	AuthorImg string
	PostType  string    `gorm:"type:varchar(100)"`
	Title     string    `gorm:"type:varchar(100)"`
	Content   string    `gorm:"type:text"`
	ImageURL  string    `gorm:"type:varchar(255)"`
	Likes     []Like    `gorm:"foreignKey:PostID"`
	Comments  []Comment `gorm:"foreignKey:PostID"`
}

func (table *PostInfo) TableName() string {
	return "post_info"
}

type PostWithLikes struct {
	PostInfo
	LikeCount    int
	UserLiked    bool
	CommentCount int // 留言數量
}

// 獲取五天內的貼文列表
func GetPostList(userID uint, postType string) ([]*PostWithLikes, error) {
	var posts []*PostWithLikes

	// 获取五天前的日期
	fiveDaysAgo := time.Now().AddDate(0, 0, -5)

	// 构建子查询
	commentCountSQL := "(SELECT count(*) FROM comments WHERE comments.post_id = post_info.id) as comment_count"
	likeCountSQL := "(SELECT count(*) FROM likes WHERE likes.post_id = post_info.id) as like_count"
	userLikedSQL := fmt.Sprintf("(SELECT count(*) > 0 FROM likes WHERE likes.post_id = post_info.id AND likes.user_id = %d) as user_liked", userID)

	// 构造查询
	query := utils.DB.Table("post_info").
		Select("post_info.*, user_basic.image_url as author_img, "+commentCountSQL+", "+likeCountSQL+", "+userLikedSQL).
		Joins("left join user_basic on post_info.author = user_basic.name").
		Where("post_info.created_at > ?", fiveDaysAgo)

	// 根据传入的 postType 添加筛选条件
	if postType != "" {
		query = query.Where("post_info.post_type = ?", postType)
	}

	err := query.Order("post_info.created_at desc").Scan(&posts).Error
	if err != nil {
		return nil, err
	}

	// 更新每个帖子的点赞状态
	// for _, post := range posts {
	// 	post.UserLiked = false
	// 	for _, like := range post.Likes {
	// 		if like.UserID == userID {
	// 			post.UserLiked = true
	// 			break
	// 		}
	// 	}
	// 	post.LikeCount = len(post.Likes)
	// }

	return posts, nil
}

// 創建貼文
func CreatePost(post *PostInfo) (*PostInfo, error) {
	// 建立新的貼文
	result := utils.DB.Create(post)
	if result.Error != nil {
		// 如果创建操作失败，返回nil和错误
		return nil, result.Error
	}
	// 如果创建操作成功，返回新创建的貼文和nil错误
	return post, nil
}

// 刪除貼文
func DeletePost(post *PostInfo) *gorm.DB {
	println(post)
	result := utils.DB.Delete(post)

	if result.Error != nil {
		println(result.Error)
		return result
	}
	return result
}

// 用戶發過的貼文數量
func GetUserPostCount(userID uint) (int64, error) {
	var count int64
	err := utils.DB.Model(&PostInfo{}).Where("author_id = ?", userID).Count(&count).Error
	return count, err
}
