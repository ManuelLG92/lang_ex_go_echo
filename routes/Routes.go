package routes

import (
	"api.go.com/echo/security"
	user "api.go.com/echo/user"
	"github.com/labstack/echo/v4"
)

func PrivateRoutes(e *echo.Echo) {
	route := e.Group("/auth")
	route.DELETE("/user/delete", user.Remove)
}

func PublicRoutes(e *echo.Echo) {
	route := e.Group("/")
	route.GET("user/:id", user.Show)
	route.POST("user/native-languages/:user-id", user.StoreUserNativeLanguages)
	route.POST("user/learning-languages/:user-id", user.StoreUserLearningLanguages)
	route.POST("user/create", user.Store)
	route.PUT("user/edit/:id", user.Update)
	route.POST("login", security.Login)
}
