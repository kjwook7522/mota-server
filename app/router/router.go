package router

import (
	"mota-server/app/server"
)

const BaseApiPath = "/api/v1"

func Init() {
	echo := server.Echo()
	baseGroup := echo.Group(BaseApiPath)

	initSentenceRouter(baseGroup)
	initRecordRouter(baseGroup)
	// write routers...
}
