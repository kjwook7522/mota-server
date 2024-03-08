package main

import (
	"mota-server/app"
	"mota-server/config"
	"mota-server/log"
	"mota-server/scheduler"
)

func main() {

	// scheduler 시스템 시작
	scheduler.Init()
	scheduler.RegisterJobs()
	scheduler.Start()

	// env 파일 로딩
	err := config.InitEnv()
	if err != nil {
		log.Error.Panic(err)
	}

	// app 시작
	motaApp := app.New()
	motaApp.ValidateGlobalInstances()
	motaApp.Start(8080)
}
