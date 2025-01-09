package main

import (
	"github.com/hashiotoko/go-sample-app/backend/config"
	"github.com/hashiotoko/go-sample-app/backend/infrastructure"
	"github.com/labstack/echo/v4"
)

func main() {
	config.LoadAppConfig()
	println("config Environment: ", config.Config.Environment)
	println("config LogLevel: ", config.Config.LogLevel)
	router := echo.New()

	infrastructure.Init(router)
	router.Start(":8888")
}
