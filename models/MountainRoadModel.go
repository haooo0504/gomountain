package models

import (
	"gomountain/utils"
	"log"

	"gorm.io/gorm"
)

type MountainRoad struct {
	gorm.Model
	Mountain string
	Road     string
	VideoUrl string
}

func (table *MountainRoad) TableName() string {
	return "mountain_road"
}

func GetMountainRoad() (map[string][]map[string]string, error) {
	var mountainRoads []MountainRoad
	result := utils.DB.Find(&mountainRoads)
	if result.Error != nil {
		log.Fatalf("查询错误: %v", result.Error)
		return nil, result.Error
	}

	// 使用映射来组织数据
	mountainMap := make(map[string][]map[string]string)
	for _, mr := range mountainRoads {
		roadMap := map[string]string{mr.Road: mr.VideoUrl}
		mountainMap[mr.Mountain] = append(mountainMap[mr.Mountain], roadMap)
	}

	return mountainMap, nil
}
