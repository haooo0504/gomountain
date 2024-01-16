package models

import (
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
	favorite := FavoriteMountainRoad{UserID: userID, MountainRoadID: mountainRoadID}
	result := utils.DB.Create(&favorite)
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
		}
		roadDetails = append(roadDetails, roadDetail)
	}

	return roadDetails, nil
}

func RemoveFromFavorites(userID uint, mountainRoadID uint) error {
	var favorite FavoriteMountainRoad

	// 查找特定的最爱记录
	result := utils.DB.Where("user_id = ? AND mountain_road_id = ?", userID, mountainRoadID).First(&favorite)
	if result.Error != nil {
		return result.Error
	}

	// 删除这条记录
	if err := utils.DB.Delete(&favorite).Error; err != nil {
		return err
	}

	return nil
}
