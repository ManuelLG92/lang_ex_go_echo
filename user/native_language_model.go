package user

import (
	"api.go.com/echo/language"
	"github.com/jinzhu/gorm"
)

type NativeLanguages struct {
	gorm.Model
	UserID     uint `json:"user_id_native_language"`
	User       User
	LanguageID uint `json:"language_id_native_language"`
	Language   language.Language
	Level      string `json:"level_native_language" validate:"required,min=2,max=10" gorm:"type:varchar(15)"`
}
