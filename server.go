package main

import (
	"api.go.com/echo/chat"
	"api.go.com/echo/config"
	"api.go.com/echo/routes"
	_ "api.go.com/echo/user"
	"fmt"
	socketio "github.com/googollee/go-socket.io"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
)

func main() {

	server := socketio.NewServer(nil)

	go chat.SetUpChat(server)

	go func() {
		err := server.Serve()
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("connected")
	}()
	defer func(server *socketio.Server) {
		err := server.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(server)

	e := echo.New()

	e.Static("/chat", "public")

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "method=${method}, uri=${uri}, status=${status}\n",
	}))

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete, http.MethodOptions},
	}))

	e.Use(middleware.Recover())

	e.Any("/socket.io/", func(c echo.Context) error {
		fmt.Println("request")
		/*		allowHeaders := "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization"

				if origin := c.Request().Header.Get("Origin"); origin != "" {
					c.Response().Header().Set("Access-Control-Allow-Origin", origin)
					c.Response().Header().Set("Vary", "Origin")
					c.Response().Header().Set("Access-Control-Allow-Methods", "POST, PUT, PATCH, GET, DELETE")
					c.Response().Header().Set("Access-Control-Allow-Credentials", "true")
					c.Response().Header().Set("Access-Control-Allow-Headers", allowHeaders)
				}
				if c.Request().Method == "OPTIONS" {
					return nil
				}
				// this takes care of the server side, No Orign no CORS baby
				c.Response().Header().Del("Origin")*/
		server.ServeHTTP(c.Response(), c.Request())
		return nil
	})
	config.InitializeDB()
	routes.PrivateRoutes(e)
	routes.PublicRoutes(e)

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.Logger.Fatal(e.Start(":1323"))
}
