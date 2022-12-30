package user

import (
	"api.go.com/echo/language"
	"github.com/jinzhu/gorm"
	"api.go.com/echo/config"

)

type NativeLanguages struct {
	gorm.Model
	UserID     uint `json:"user_id"`
	User       User
	LanguageID uint `json:"language_id"`
	Language   language.Language
	Level      string `json:"level" validate:"required,min=2,max=10" gorm:"type:varchar(15)"`
}

func MigrateLearningNativeLanguagesTable() {
	config.DbGlobal.AutoMigrate(&NativeLanguages{})
}
