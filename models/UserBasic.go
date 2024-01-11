package models

import (
	"gomountain/utils"
	"time"

	"gorm.io/gorm"
)

type UserBasic struct {
	gorm.Model
	Name          string `gorm:"not null"`
	Password      string
	Phone         string
	Email         string `gorm:"not null;valid:\"email\""`
	Identity      string
	ClientIp      string
	ClientPort    string
	Salt          string
	LoginTime     uint64
	HeartbeatTime uint64
	LogOutTime    uint64
	IsLogout      bool
	DeviceInfo    string
	CanUseTime    uint64
	ImageURL      string
	IsGoogle      bool
	IsApple       bool
	TokenSub      string
	Likes         []Like `gorm:"foreignKey:UserID"`
}

type UserAccountInfo struct {
	ID        uint      // 用户ID通常是账号的唯一标识
	CreatedAt time.Time // GORM Model的CreatedAt字段用于存储记录的创建时间
	Name      string
	Email     string
	IsGoogle  bool
	IsApple   bool
	ImageURL  string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}

func GetUserList() ([]*UserAccountInfo, error) {
	var data []*UserAccountInfo
	result := utils.DB.Model(&UserBasic{}).Select("id", "created_at", "name", "email", "is_google", "is_apple", "image_url").Find(&data)
	if result.Error != nil {
		return nil, result.Error
	}
	return data, nil
}

func FindUserByNameAndPwd(name string, password string, token string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ? and password = ?", name, password).First(&user)
	// token加密
	// str := fmt.Sprintf("%d", time.Now().Unix())
	// temp := utils.MD5Encode(str)
	utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", token)
	return user
}

func FindUserByName(name string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("name = ?", name).First(&user)
	return user
}

func FindUserById(id uint) UserBasic {
	user := UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	return user
}

func FindUserByPhone(phone string) *gorm.DB {
	user := UserBasic{}
	return utils.DB.Where("phone = ?", phone).First(&user)
}

func FindUserByEmail(email string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("email = ?", email).First(&user)
	return user
}

func CreateUser(user *UserBasic) (*UserBasic, error) {
	result := utils.DB.Create(user)
	if result.Error != nil {
		// 如果创建操作失败，返回nil和错误
		return nil, result.Error
	}
	// 如果创建操作成功，返回新创建的用户和nil错误
	return user, nil
}

func DeleteUser(user UserBasic) *gorm.DB {
	return utils.DB.Delete(&user)
}

func UpdateUser(user UserBasic) (UserBasic, error) {
	if err := utils.DB.Model(&user).Updates(UserBasic{Name: user.Name, Password: user.Password, Phone: user.Phone, Email: user.Email, ImageURL: user.ImageURL, Salt: user.Salt, Identity: user.Identity}).Error; err != nil {
		return UserBasic{}, err
	}

	// 重新查詢更新後的用戶數據
	var updatedUser UserBasic
	if err := utils.DB.First(&updatedUser, user.ID).Error; err != nil {
		return UserBasic{}, err
	}

	return updatedUser, nil
}

func RefreshToken(id uint, name string, oldToken string, token string) (UserBasic, bool) {
	user := UserBasic{}
	utils.DB.Where("id = ?", id).First(&user)
	if user.Name == name && user.Identity == oldToken {
		utils.DB.Model(&user).Where("id = ?", user.ID).Update("identity", token)
		return user, true
	} else {
		return user, false
	}
}

func FindUserByGoogleSignIn(email string, sub string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("email = ? AND is_google = ? AND token_sub = ?", email, true, sub).First(&user)
	return user
}

func FindUserByAppleSignIn(email string, sub string) UserBasic {
	user := UserBasic{}
	utils.DB.Where("email = ? AND is_apple = ? AND token_sub = ?", email, true, sub).First(&user)
	return user
}
