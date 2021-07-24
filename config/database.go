package config

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)
var DbGlobal *gorm.DB

func InitializeDB(){
	dsn := "root:@tcp(127.0.0.1:3306)/golang?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		fmt.Println("DB conection error.")
		panic(err)
	}
		DbGlobal = db
		fmt.Println("Connection successfully to database.")
}
