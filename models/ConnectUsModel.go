package models

import (
	"gomountain/utils"

	"gorm.io/gorm"
)

type ConnectUs struct {
	gorm.Model
	UserID   string // 用戶ID
	UserName string // 用戶名字
	Gmail    string // 電子郵件
	Content  string `gorm:"type:text"`
}

func (ConnectUs) TableName() string {
	return "connect_us"
}

func ConnectToUs(userID string, userName string, gmail string, content string) (*ConnectUs, error) {
	comment := ConnectUs{
		UserID:   userID,
		UserName: userName,
		Gmail:    gmail,
		Content:  content,
	}

	if err := utils.DB.Create(&comment).Error; err != nil {
		return nil, err
	}

	return &comment, nil
}
