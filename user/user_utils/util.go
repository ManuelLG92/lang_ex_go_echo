package user_utils

import (
	"api.go.com/echo/config"
	"api.go.com/echo/country"
	"api.go.com/echo/language"
	"api.go.com/echo/user"
)


func CreateTables() {
	config.DbGlobal.AutoMigrate(
		&language.Language{}, &country.Country{},
		&user.Gender{}, &user.User{},
		&user.LearningLanguages{}, &user.NativeLanguages{})
}
