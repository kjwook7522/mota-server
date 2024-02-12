package router

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Router struct {
	echo *echo.Echo
	base *echo.Group
	db   *gorm.DB
}

const BASE_API_PATH = "/api/v1"

func New(e *echo.Echo, db *gorm.DB) *Router {
	router := Router{echo: e, db: db}
	router.base = e.Group(BASE_API_PATH)

	return &router
}

func (r *Router) RegisterAll() {
}
