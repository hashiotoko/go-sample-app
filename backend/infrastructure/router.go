package infrastructure

import (
	"github.com/labstack/echo/v4"
)

func Init(router *echo.Echo) {
  router.GET("/", func(c echo.Context) error {
    return c.String(200, "Hello, World!")
  })
}
