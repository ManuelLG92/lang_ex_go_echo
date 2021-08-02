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

	e.Any("/socket.io/", func(context echo.Context) error {
		fmt.Println("request")
		server.ServeHTTP(context.Response(), context.Request())
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
