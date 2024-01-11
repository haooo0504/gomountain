package models

import (
	"gomountain/utils"

	"gorm.io/gorm"
)

type Comment struct {
	gorm.Model
	UserID       uint   // 用户ID
	PostID       uint   // 帖子ID
	Content      string `gorm:"type:text"` // 评论内容
	UserName     string `gorm:"-"`         // 非数据库字段，用于存储用户的名字
	UserImageUrl string `gorm:"-"`         // 非数据库字段，用于存储用户的名字
	// 可以添加其他字段，如时间戳等
}

type CommentWithUser struct {
	gorm.Model
	UserID       uint
	PostID       uint
	Content      string
	UserName     string `gorm:"column:user_name"`      // 确保这里的 column 名称和 SQL 查询的别名一致
	UserImageUrl string `gorm:"column:user_image_url"` // 确保这里的 column 名称和 SQL 查询的别名一致
}

func (Comment) TableName() string {
	return "comments"
}

func AddComment(userID uint, postID uint, content string) (*Comment, error) {
	comment := Comment{
		UserID:  userID,
		PostID:  postID,
		Content: content,
	}

	if err := utils.DB.Create(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}

func GetCommentsByPostID(postID uint) ([]CommentWithUser, error) {
	var comments []CommentWithUser
	err := utils.DB.Table("comments").
		Select("comments.*, user_basic.name as user_name, user_basic.image_url as user_image_url").
		Joins("left join user_basic on user_basic.id = comments.user_id").
		Where("comments.post_id = ?", postID).
		Scan(&comments).Error
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// func GetCommentsByPostID(postID uint) ([]Comment, error) {
//     var comments []Comment
//     sql := "SELECT comments.*, user_basic.name as user_name FROM comments LEFT JOIN user_basic ON user_basic.id = comments.user_id WHERE comments.post_id = ?"
//     err := utils.DB.Raw(sql, postID).Scan(&comments).Error
//     if err != nil {
//         return nil, err
//     }

//     return comments, nil
// }
