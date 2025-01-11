package interfaces

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

type GreetingController struct{}

func NewGreetingController() *GreetingController {
	return &GreetingController{}
}

func (c *GreetingController) GreetingGreeting(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, "Hello, world!")
}
