package app

import (
	"mota-server/app/router"
	"mota-server/app/server"
	"mota-server/log"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MotaApp struct {
	echo *echo.Echo
	db   *gorm.DB
}

func New() *MotaApp {

	// db 로딩
	err := server.InitDB()
	if err != nil {
		log.Error.Panic(err)
	}

	// echo 로딩
	server.InitEcho()
	server.InitEchoMiddleware()

	// router 로딩
	router.Init()

	// app 생성
	app := &MotaApp{echo: server.Echo(), db: server.DB()}

	return app
}

func (app *MotaApp) Start(port int) {
	app.echo.Logger.Fatal(app.echo.Start(":" + strconv.Itoa(port)))
}

func (app *MotaApp) ValidateGlobalInstances() {
	if app.echo == nil {
		log.Error.Panic("echo instance has not been initialized")
	}

	if app.db == nil {
		log.Error.Panic("db instance has not been initialized")
	}
}
