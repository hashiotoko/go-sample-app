package main

import (
	"github.com/hashiotoko/go-sample-app/backend/infrastructure"
	"github.com/labstack/echo/v4"
)

func main() {
	router := echo.New()

	infrastructure.Init(router)
	router.Start(":8888")
}
