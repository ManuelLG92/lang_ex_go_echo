package user

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strconv"
)

func StoreUserNativeLanguages(c echo.Context) error {

	id := c.Param("user-id")

	err, user := checkIfUserExistById(id)
	if err != nil {
		return err
	}

	var nativeLanguages []*NativeLanguages

	err = c.Bind(&nativeLanguages)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	for _, language := range nativeLanguages {

		if language.UserID != user.ID {
			return echo.NewHTTPError(http.StatusBadRequest, "User_id in native language doesn't match,")
		}

		if err := checkIfLanguageIdExists(language.LanguageID); err != nil {
			return echo.NewHTTPError(http.StatusNotFound, "Language "+strconv.Itoa(int(language.LanguageID))+" not found")
		}

	}

	if err := InsertIntoDbNativeLanguagesFromArray(nativeLanguages); err != nil {
		return err
	}

	return nil
}
