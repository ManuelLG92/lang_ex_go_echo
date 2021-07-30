package routes

import (
	"api.go.com/echo/language"
	"api.go.com/echo/security"
	"api.go.com/echo/user"
	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"strconv"
)

func PrivateRoutes(e *echo.Echo) {
	route := e.Group("/auth/")
	config := middleware.JWTConfig{
		Claims:     &security.JwtCustomClaims{},
		SigningKey: []byte("secret"),
		ContextKey: "authenticated",
	}
	route.Use(middleware.JWTWithConfig(config))
	route.GET("user/:id", user.Show)
	route.PUT("user/edit/:id", user.Update, checkIfUserIdBelongsTAuthenticatedUser)
	route.DELETE("/user/delete", user.Remove)
}

func PublicRoutes(e *echo.Echo) {
	route := e.Group("/")
	route.POST("user/native-languages/:user-id", user.StoreUserNativeLanguages)
	route.POST("user/learning-languages/:user-id", user.StoreUserLearningLanguages)
	route.POST("user/create", user.Store)
	route.GET("languages", language.GetLanguages)
	route.POST("login", security.Login)
	user.MigrateUserTable()
}

func checkIfUserIdBelongsTAuthenticatedUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			return c.JSON(http.StatusBadRequest, err)
		}
		userJWT := c.Get("authenticated").(*jwt.Token)
		claims := userJWT.Claims.(*security.JwtCustomClaims)

		if uint(id) != claims.Id {
			return echo.NewHTTPError(http.StatusUnauthorized, "You only can update your own profile.")
		}
		return next(c)
	}
}
