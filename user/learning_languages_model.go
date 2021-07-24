package user

import (
	"github.com/jinzhu/gorm"
)

type LearningLanguages struct {
	gorm.Model
	UserID     uint `json:"user_id_learning_languages"`
	User       User
	LanguageID uint   `json:"language_id_learning_languages"`
	Level      string `json:"level_native_language" validate:"required,min=2,max=10" gorm:"type:varchar(15)"`
}
