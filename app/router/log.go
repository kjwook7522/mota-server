package router

import (
	"github.com/labstack/echo/v4"
	"mota-server/app/controller"
	"mota-server/app/repository"
	"mota-server/app/server"
)

func initLogRouter(base *echo.Group) {
	var route *echo.Route
	db := server.DB()

	shortSentencePlayLogRepo := repository.NewShortSentencePlayLogRepository(db)
	ctr := controller.NewLogController(shortSentencePlayLogRepo)

	route = base.GET("/log/play/sentence/short", ctr.GetShortSentencePlayLogs)
	route.Name = "get short sentence play logs"

	route = base.POST("/log/play/sentence/short", ctr.CreateShortSentencePlayLogs)
	route.Name = "create short sentence play logs"
}
