package server

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"mota-server/log"
)

var e *echo.Echo

func InitEcho() {
	e = echo.New()
}

func Echo() *echo.Echo {
	return e
}

func InitEchoMiddleware() {
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format:           "[API]${time_custom} ${method} ${status} ${uri} ${remote_ip} ${user_agent} ${error}\n",
		CustomTimeFormat: "2006/01/02 15:04:05",
		Output:           log.File(),
	}))
}
