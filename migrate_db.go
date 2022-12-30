package main

import (
	"api.go.com/echo/user"
)

func migrate()  {
	user.MigrateUserTable()
	user.MigrateGenderTable()
	user.MigrateLearningLanguagesTable()
	user.MigrateLearningNativeLanguagesTable()

}