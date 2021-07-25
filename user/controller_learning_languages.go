package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func StoreUserLearningLanguages(c echo.Context) error {

	id := c.Param("user-id")

	err, user := checkIfUserExistById(id)
	if err != nil {
		return err
	}

	var learningLanguages []*LearningLanguages

	err = c.Bind(&learningLanguages)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, language := range learningLanguages {

		if language.UserID != user.ID {
			return echo.NewHTTPError(http.StatusBadRequest, "User_id in native language doesn't match,")
		}

		if err := checkIfLanguageIdExists(language.LanguageID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Language "+strconv.Itoa(int(language.LanguageID))+" not found")
		}

	}

	if err := InsertIntoDbLearningLanguagesFromArray(learningLanguages); err != nil {
		return err
	}

	return nil
}
