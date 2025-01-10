package usecases

import (
	"context"

	"github.com/hashiotoko/go-sample-app/backend/domain"
)

type UserRepository interface {
	GetUsers(ctx context.Context) ([]domain.User, error)
	GetUsersByID(ctx context.Context, id string) (domain.User, error)
}
