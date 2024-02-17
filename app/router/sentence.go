package router

import (
	"github.com/labstack/echo/v4"
	"mota-server/app/controller"
	"mota-server/app/repository"
	"mota-server/app/server"
)

func initSentenceRouter(base *echo.Group) {
	var route *echo.Route
	db := server.DB()

	shortSentenceRepo := repository.NewShortSentenceRepository(db)
	ctr := controller.NewShortSentenceController(shortSentenceRepo)

	route = base.GET("/sentence/short", ctr.GetShortSentences)
	route.Name = "get short sentences"
}
