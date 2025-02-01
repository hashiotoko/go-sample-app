package usecases

import (
	"context"

	"github.com/hashiotoko/go-sample-app/backend/domain"
	"github.com/hashiotoko/go-sample-app/backend/usecases/repository_interface/dto"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUsersByID(ctx context.Context, id string) (domain.User, error)
	CreateUser(ctx context.Context, req dto.CreateUserRequest) (domain.User, error)
}
