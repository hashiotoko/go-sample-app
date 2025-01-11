package interfaces

import (
	"net/http"
	"strconv"

	api "github.com/hashiotoko/go-sample-app/backend/api/generated"
	"github.com/hashiotoko/go-sample-app/backend/infrastructure/db"
	repositories "github.com/hashiotoko/go-sample-app/backend/interfaces/repositories"
	"github.com/hashiotoko/go-sample-app/backend/usecases"
	"github.com/hashiotoko/go-sample-app/backend/usecases/dto"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	Interactor usecases.UserInteractor
}

func NewUserController(dbClinet db.Client) *UserController {
	interactor := usecases.NewUserInteractor(
		repositories.NewUserRepository(dbClinet),
	)
	return &UserController{
		Interactor: interactor,
	}
}

func (c *UserController) UsersGetUsers(ctx echo.Context) error {
	dtos, err := c.Interactor.GetUsers(ctx.Request().Context())
	if err != nil {
		return err
	}
	users := make([]api.ModelsUser, 0)
	for _, dto := range dtos {
		tmp := convertUser(dto)
		users = append(users, tmp)
	}

	return ctx.JSON(http.StatusOK, users)
}

func (c *UserController) UsersGetUser(ctx echo.Context, userId int32) error {
	id := strconv.Itoa(int(userId))

	dto, err := c.Interactor.GetUsersByID(ctx.Request().Context(), id)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, convertUser(dto))
}

func convertUser(u dto.User) api.ModelsUser {
	return api.ModelsUser{
		Id:   u.ID,
		Name: u.Name,
	}
}
