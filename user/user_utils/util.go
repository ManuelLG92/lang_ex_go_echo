package user_utils

import (
	"github.com/ManuelLG92/lang_ex_api/config"
	"github.com/ManuelLG92/lang_ex_api/country"
	"github.com/ManuelLG92/lang_ex_api/language"
	"github.com/ManuelLG92/lang_ex_api/user"
)

func CreateTables() {
	config.DbGlobal.AutoMigrate(
		&language.Language{}, &country.Country{},
		&user.Gender{}, &user.User{},
		&user.LearningLanguages{}, &user.NativeLanguages{})
}
