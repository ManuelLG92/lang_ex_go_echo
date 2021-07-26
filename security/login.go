package security

import (
	"api.go.com/echo/config"
	"api.go.com/echo/globals/password"
	"api.go.com/echo/user"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"
)

type LoginStruct struct {
	Email    string `json:"email" validate:"email"`
	Password string `json:"password" validate:"required,min=8,max=32"`
}

func Login(c echo.Context) error {

	var login LoginStruct

	if err := c.Bind(&login); err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Couldn't get user credentials")
	}

	var userInstance user.User

	if err := config.DbGlobal.Where("email = ?", login.Email).First(&userInstance); err.Error != nil {
		fmt.Println(userInstance)
		return echo.NewHTTPError(http.StatusNotFound, "User not found for email %s", login.Email)
	}

	if err := password.Verify(userInstance.Password, login.Password); err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, "Invalid password.")
	}

	claims := &JwtCustomClaims{
		userInstance.ID,
		userInstance.Name,
		c.RealIP(),
		checkBooleanRole(userInstance.IsModerator),
		checkBooleanRole(userInstance.IsAdmin),
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 20).Unix(),
		},
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	signedStringToken, err := token.SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	return c.JSON(http.StatusAccepted, echo.Map{
		"token": signedStringToken,
	})

}

func checkBooleanRole(role uint8) bool {

	if role == 1 {
		return true
	}
	return false
}

/*type ErrorResponse struct {
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
*/
