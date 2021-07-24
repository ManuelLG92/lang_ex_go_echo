package user

import "github.com/jinzhu/gorm"

type Gender struct {
	gorm.Model
	Name string `gorm:"type:varchar(20)"`
}
