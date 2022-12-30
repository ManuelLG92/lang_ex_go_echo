package user

import (
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name" validate:"required,min=2,max=150" gorm:"type:varchar(150)"`
	Surname     string `json:"surname"  validate:"required,min=2,max=150" gorm:"type:varchar(150)"`
	Email       string `json:"email" validate:"required,email" gorm:"unique"`
	Password    string `json:"password" validate:"required,min=8"`
	Birthday    string `json:"birthday"`
	Country   string    `json:"country_id" validate:"required"`
	Gender    string    `json:"gender_id"  validate:"required"`
	Description string `json:"description" gorm:"type:text" validate:"max=255"`
	IsAdmin     uint8  `gorm:"default:0"`
	IsModerator uint8  `gorm:"default:0"`
}

type DTOUser struct {
	ID          uint             `json:"id" validate:"required,min=1"`
	Name        string           `json:"name" validate:"required,min=2,max=150"`
	Surname     string           `json:"surname"  validate:"required,min=2,max=150"`
	Birthday    string           `json:"birthday"`
	Country     string          `json:"country_id" validate:"required"`
	Gender      string          `json:"gender_id"  validate:"required"`
	Description string           `json:"description" gorm:"type:text" validate:"max=255"`
}
