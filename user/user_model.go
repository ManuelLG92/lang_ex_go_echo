package user

import (
	"api.go.com/echo/country"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name        string `json:"name" validate:"required,min=2,max=150" gorm:"type:varchar(150)"`
	Surname     string `json:"surname"  validate:"required,min=2,max=150" gorm:"type:varchar(150)"`
	Email       string `json:"email" validate:"required,email" gorm:"unique"`
	Password    string `json:"password" validate:"required,min=8"`
	Birthday    string `json:"birthday"`
	CountryID   int    `json:"country_id" validate:"required"`
	GenderID    int    `json:"gender_id"  validate:"required"`
	Description string `json:"description" gorm:"type:text" validate:"max=255"`
	IsAdmin     uint8  `gorm:"default:0"`
	IsModerator uint8  `gorm:"default:0"`
}

type DTOUser struct {
	ID          uint             `json:"user_id" validate:"required,min=1"`
	Name        string           `json:"user_name" validate:"required,min=2,max=150"`
	Surname     string           `json:"user_surname"  validate:"required,min=2,max=150"`
	Birthday    string           `json:"user_birthday"`
	Country     *country.Country `json:"user_country_id" validate:"required"`
	Gender      *Gender          `json:"user_gender_id"  validate:"required"`
	Description string           `json:"user_description" gorm:"type:text" validate:"max=255"`
}
