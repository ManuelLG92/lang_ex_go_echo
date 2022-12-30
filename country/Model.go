package country

import (
	"api.go.com/echo/config"
	"github.com/jinzhu/gorm"
)

type Country struct {
	gorm.Model
	Name    string `json:"name" gorm:"type:varchar(50)"`
	IsoCode string `json:"iso_code" gorm:"type:varchar(10)"`
}

func MigrateCountryTable() {
	config.DbGlobal.AutoMigrate(&Country{})
}
