package main

import (
	"github.com/labstack/echo"
)

func initServer(e *echo.Echo) {
	e.Start("0.0.0.0:2327")
}
