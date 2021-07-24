package userdomain

import (
	"api.go.com/echo/config"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID           uint 		`json:"id"`
	Name         string ` sql:"type:VARCHAR(100)" json:"name"`
	Email        string  `sql:"type:VARCHAR(150)" gorm:"unique" json:"email"`
	Password 	 string   `json:"password"`
}

func CreateUserDB()  {
	err := config.DbGlobal.AutoMigrate(&User{})
	if err != nil {
		panic(err)
	}
}


