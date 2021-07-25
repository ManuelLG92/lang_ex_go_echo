package user

import (
	"api.go.com/echo/config"
	"github.com/labstack/echo/v4"
	"net/http"
)

func InsertIntoDbLearningLanguagesFromArray(learningLanguageArray []*LearningLanguages) error {
	for _, singleLearningLang := range learningLanguageArray {
		if err := config.DbGlobal.Create(&singleLearningLang); err.Error != nil {
			return echo.NewHTTPError(http.StatusNotImplemented, err.Error)
		}
	}
	return nil
}
