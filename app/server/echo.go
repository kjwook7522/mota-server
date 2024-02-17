package server

import (
	"github.com/labstack/echo/v4"
)

var e *echo.Echo

func InitEcho() {
	e = echo.New()
}

func Echo() *echo.Echo {
	return e
}
