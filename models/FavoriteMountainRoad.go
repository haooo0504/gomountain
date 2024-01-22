package models

import (
	"errors"
	"gomountain/utils"
	"strconv"

	"gorm.io/gorm"
)

type FavoriteMountainRoad struct {
	gorm.Model
	UserID         uint
	MountainRoadID uint
}

func (table *FavoriteMountainRoad) TableName() string {
	return "favorite_mountain_road"
}

func AddToFavorites(userID uint, mountainRoadID uint) error {
	var existingFavorite FavoriteMountainRoad
	// 檢查是否已經存在相同的收藏紀錄
	result := utils.DB.Where("user_id = ? AND mountain_road_id = ?", userID, mountainRoadID).First(&existingFavorite)

	// 如果紀錄已經存在則拋出錯誤
	if result.RowsAffected > 0 {
		return errors.New("已經加入最愛")
	}
	favorite := FavoriteMountainRoad{UserID: userID, MountainRoadID: mountainRoadID}
	result = utils.DB.Create(&favorite)
	return result.Error
}

func GetUserFavoriteMountainRoads(userID uint) ([]map[string]string, error) {
	var favoriteRoads []FavoriteMountainRoad
	var roadDetails []map[string]string

	// 查找用户的所有最爱 MountainRoad
	result := utils.DB.Where("user_id = ?", userID).Find(&favoriteRoads)
	if result.Error != nil {
		return nil, result.Error
	}

	// 对于每个最爱的 MountainRoad，获取其 Mountain, Road 和 VideoUrl
	for _, fav := range favoriteRoads {
		var mountainRoad MountainRoad
		if err := utils.DB.First(&mountainRoad, fav.MountainRoadID).Error; err != nil {
			return nil, err
		}
		roadDetail := map[string]string{
			"id":       strconv.FormatUint(uint64(mountainRoad.ID), 10),
			"mountain": mountainRoad.Mountain,
			"road":     mountainRoad.Road,
			"videoUrl": mountainRoad.VideoUrl,
			"city":     mountainRoad.City,
			"area":     mountainRoad.Area,
		}
		roadDetails = append(roadDetails, roadDetail)
	}

	return roadDetails, nil
}

func RemoveFromFavorites(userID uint, mountainRoadID uint) error {
	var favorite FavoriteMountainRoad

	// 查找指定紀錄
	result := utils.DB.Where("user_id = ? AND mountain_road_id = ?", userID, mountainRoadID).First(&favorite)
	if result.Error != nil {
		return result.Error
	}

	// 執行硬刪除(真實刪除)這條紀錄
	if err := utils.DB.Unscoped().Delete(&favorite).Error; err != nil {
		return err
	}

	return nil
}
