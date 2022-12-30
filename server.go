package main

import (
	"api.go.com/echo/config"
	"api.go.com/echo/routes"
	"api.go.com/echo/user"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

type responseWs struct {
	socketId string
	username string
}

func main() {

	e := echo.New()

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	e.Use(middleware.Recover())
	config.InitializeDB()
	user.MigrateUserTable()
	user.MigrateGenderTable()
	user.MigrateLearningLanguagesTable()
	user.MigrateLearningNativeLanguagesTable()
	routes.PrivateRoutes(e)
	routes.PublicRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
