package user

import (
	"api.go.com/echo/config"
)

func InsertIntoDbNativeLanguagesFromArray(nativeLanguagesArray []*NativeLanguages) {
	for _, singleNativeLang := range nativeLanguagesArray {
		config.DbGlobal.Create(&singleNativeLang)
	}
}
