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

const BaseApiPath = "/api/v1"

func New(e *echo.Echo, db *gorm.DB) *Router {
	router := Router{echo: e, db: db}
	router.base = e.Group(BaseApiPath)

	return &router
}

func (r *Router) RegisterAll() {
	r.initLogRouter()
}
