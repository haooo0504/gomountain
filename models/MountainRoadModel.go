package models

import (
	"gomountain/utils"
	"log"
	"strconv"

	"gorm.io/gorm"
)

type MountainRoad struct {
	gorm.Model
	Mountain string
	Road     string
	VideoUrl string
	City     string
	Area     string
}

func (table *MountainRoad) TableName() string {
	return "mountain_road"
}

func GetMountainRoad() (map[string][]map[string]string, error) {
	var mountainRoads []MountainRoad
	result := utils.DB.Find(&mountainRoads)
	if result.Error != nil {
		log.Fatalf("查詢錯誤: %v", result.Error)
		return nil, result.Error
	}

	// 使用映射来组织数据
	mountainMap := make(map[string][]map[string]string)
	for _, mr := range mountainRoads {
		roadMap := map[string]string{"id": strconv.FormatUint(uint64(mr.ID), 10), "road": mr.Road, "videoUrl": mr.VideoUrl, "city": mr.City, "area": mr.Area}
		mountainMap[mr.Mountain] = append(mountainMap[mr.Mountain], roadMap)
	}

	return mountainMap, nil
}
