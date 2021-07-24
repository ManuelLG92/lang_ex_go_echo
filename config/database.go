package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DbGlobal *gorm.DB

func InitializeDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/lang_ex?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err.Error())
	}

	DbGlobal = db

	fmt.Println("Database connected")

	return DbGlobal
}
