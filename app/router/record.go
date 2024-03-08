package router

import (
	"github.com/labstack/echo/v4"
	"mota-server/app/controller"
	"mota-server/app/repository"
	"mota-server/app/server"
)

func initRecordRouter(base *echo.Group) {
	var route *echo.Route
	db := server.DB()

	shortSentenceRecordRepo := repository.NewShortSentenceRecordRepository(db)
	ctr := controller.NewRecordController(shortSentenceRecordRepo)

	route = base.GET("/record/sentence/short", ctr.GetShortSentenceRecords)
	route.Name = "get short sentence records"

	route = base.POST("/record/sentence/short", ctr.CreateShortSentenceRecords)
	route.Name = "create short sentence records"
}
