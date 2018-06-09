package controller

import (
	"net/http"

	"github.com/labstack/echo"
)

// InitController initialize controller
func InitController(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error { return c.String(http.StatusOK, "inandout API is Alive!!") })
}
