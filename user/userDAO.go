package user

import (
	"api.go.com/echo/config"
	"gorm.io/gorm"
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

var userGlobal *User

func MigrateUserTable() {
	config.DbGlobal.AutoMigrate(&User{})
}


func (user User) CreateUser() *gorm.DB {
	return config.DbGlobal.Create(&user)
}

func (user User) UpdateUser() *gorm.DB {
	return config.DbGlobal.Save(&user)
}

func GetUser(userId int) (error,*User) {
	value := config.DbGlobal.First(&userGlobal, userId)
	if value.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "User %v not found", strconv.Itoa(int(userId))), nil
	}
	return nil, userGlobal
}

func (user User) DeleteUser() *gorm.DB {
	return config.DbGlobal.Delete(&User{}, user.ID)
}

func (user User) GetAllNativeLanguages() (error, []*NativeLanguages) {
	var nativeLanguages []*NativeLanguages
	if err := config.DbGlobal.Where("user_id  = ?", user.ID).Find(&nativeLanguages); err.Error != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't retrieve native languages for user %v", strconv.Itoa(int(user.ID))), nativeLanguages
	}

	return nil, nativeLanguages
}

func (user User) GetAllLearningLanguages() (error, []*LearningLanguages) {
	var learningLanguages []*LearningLanguages
	if err := config.DbGlobal.Where("user_id  = ?", user.ID).Find(&learningLanguages); err != nil {
		return err.Error, learningLanguages
	}
	return nil, learningLanguages

}
