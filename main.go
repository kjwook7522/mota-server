package main

import (
	"log"
	"mota-server/app"
	"mota-server/config"
)

func main() {
	// env 파일 로딩
	err := config.InitEnv()
	if err != nil {
		log.Fatal(err)
	}

	// app 초기 시작
	motaApp := app.New()
	motaApp.Start(8080)
}
