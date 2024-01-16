package main

import (
	"gomountain/models"
	"gomountain/router"
	"gomountain/utils"
	"log"
)

// @securityDefinitions.apikey	ApiKeyAuth
// @in							header
// @name						Authorization
// @description				Description for what is this security definition being used
func main() {
	utils.InitConfig()
	utils.InitMySQL()

	// Add AutoMigrate calls here
	// 進行數據庫遷移
	db := utils.GetDB()
	err := db.AutoMigrate(&models.UserBasic{})
	if err != nil {
		log.Fatalf("failed to auto migrate UserBasic: %v", err)
	}
	err = db.AutoMigrate(&models.PostInfo{})
	if err != nil {
		log.Fatalf("failed to auto migrate PostInfo: %v", err)
	}
	err = db.AutoMigrate(&models.Like{})
	if err != nil {
		log.Fatalf("failed to auto migrate Like: %v", err)
	}
	err = db.AutoMigrate(&models.Comment{})
	if err != nil {
		log.Fatal("failed to auto migrate Comment: ", err)
	}
	err = db.AutoMigrate(&models.Version{})
	if err != nil {
		log.Fatal("failed to auto migrate Version: ", err)
	}
	err = db.AutoMigrate(&models.MountainRoad{})
	if err != nil {
		log.Fatal("failed to auto migrate MountainRoad: ", err)
	}

	r := router.Router()
	r.Run(":8083")
}
