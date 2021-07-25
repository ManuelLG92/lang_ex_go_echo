package user

import (
	"api.go.com/echo/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InsertIntoDbNativeLanguagesFromArray(nativeLanguagesArray []*NativeLanguages) error {
	for _, singleNativeLang := range nativeLanguagesArray {
		if err := config.DbGlobal.Create(&singleNativeLang); err.Error != nil {
			return echo.NewHTTPError(http.StatusNotImplemented, err.Error)
		}
	}

	return nil
}
