package app

import (
	"log"
	"mota-server/db"
	"mota-server/router"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type MotaApp struct {
	echo   *echo.Echo
	db     *gorm.DB
	router *router.Router
}

func New() *MotaApp {

	// db 로딩
	d, err := db.New()
	if err != nil {
		log.Fatal(err)
	}

	// echo 로딩
	e := echo.New()

	// router 로딩
	r := router.New(e, d)
	r.RegisterAll()

	// app 생성
	app := &MotaApp{echo: e, db: d, router: r}

	return app
}

func (app *MotaApp) Start(port int) {
	app.echo.Logger.Fatal(app.echo.Start(":" + strconv.Itoa(port)))
}
