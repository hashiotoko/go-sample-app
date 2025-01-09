package interfaces

import (
	"net/http"

	api "github.com/hashiotoko/go-sample-app/backend/api/generated"
	"github.com/labstack/echo/v4"
)

type UserController struct{}

var users = []api.User{
	{
		Id:   1,
		Name: "田中太郎",
	},
	{
		Id:   2,
		Name: "山田次郎",
	},
}

func NewUserController() *UserController {
	return &UserController{}
}

func (c *UserController) GetApiV1Users(ctx echo.Context) error {
	return ctx.JSON(http.StatusOK, users)
}

func (c *UserController) GetApiV1UsersUserId(ctx echo.Context, userId int32) error {
	for _, user := range users {
		if user.Id == userId {
			return ctx.JSON(http.StatusOK, user)
		}
	}

	return ctx.JSON(http.StatusNotFound, map[string]string{
		"message": "user not found",
	})
}
