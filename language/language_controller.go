package language

import (
	"api.go.com/echo/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

func GetLanguages(c echo.Context) error {
	var languages []Language
	if err := config.DbGlobal.Find(&languages); err.Error != nil {
		return c.JSON(http.StatusNotFound, err.Error)
	}

	return c.JSON(http.StatusAccepted, languages)
}
