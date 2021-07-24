package security

import (
	"api.go.com/echo/config"
	"api.go.com/echo/globals/password"
	"api.go.com/echo/user"
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"net/http"
)

type LoginStruct struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func Login(c echo.Context) error {

	login := new(LoginStruct)

	if err := c.Bind(&login); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't get user credentials")
	}

	userInstance := new(user.User)
	fmt.Println(login.Email)
	if err := config.DbGlobal.Where("email = ?", login.Email).First(&userInstance); err.Error != nil {
		fmt.Println(userInstance)
		return echo.NewHTTPError(http.StatusNotFound, "User not found for email %s", login.Email)
	}

	if err := password.Verify(userInstance.Password, login.Password); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid password.")
	}

	return nil
	//return nil

}

type ErrorResponse struct {
	FailedField string
	Tag         string
	Value       string
}

func ValidateStruct(user LoginStruct) []*ErrorResponse {
	var errors []*ErrorResponse
	validate := validator.New()
	err := validate.Struct(user)
	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var element ErrorResponse
			element.FailedField = err.StructNamespace()
			element.Tag = err.Tag()
			element.Value = err.Param()
			errors = append(errors, &element)
		}
	}
	return errors
}
