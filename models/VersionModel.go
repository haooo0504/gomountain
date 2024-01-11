package models

import (
	"fmt"
	"gomountain/utils"

	"gorm.io/gorm"
)

type Version struct {
	gorm.Model
	CurVersion string
}

func (table *Version) TableName() string {
	return "version"
}

func GetVersion() (*Version, error) {
	var version Version
	result := utils.DB.Model(&Version{}).Order("id DESC").First(&version) // 獲取最新版本記錄
	if result.Error != nil {
		return nil, fmt.Errorf("failed to query latest version: %w", result.Error)
	}
	return &version, nil
}
