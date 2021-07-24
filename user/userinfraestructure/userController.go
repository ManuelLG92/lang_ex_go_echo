package userinfraestructure

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func UserRoutes(e *echo.Echo) {
	e.GET("/users/:id", getUser)
}

func getUser(c echo.Context) error {
	id := c.Param("id")
	return c.String(http.StatusOK, id)
}
