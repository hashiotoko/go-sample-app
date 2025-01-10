package usecases

import (
	"context"
	"log/slog"

	"github.com/hashiotoko/go-sample-app/backend/domain"
	"github.com/hashiotoko/go-sample-app/backend/usecases/dto"
	repository "github.com/hashiotoko/go-sample-app/backend/usecases/repository_interface"
)

type userInteractor struct {
	UserRepository repository.UserRepository
}

type UserInteractor interface {
	GetUsers(ctx context.Context) ([]dto.User, error)
	GetUsersByID(ctx context.Context, id string) (dto.User, error)
}

// var _ UserInteractor = &userInteractor{}

func NewUserInteractor(userRepository repository.UserRepository) UserInteractor {
	return &userInteractor{
		UserRepository: userRepository,
	}
}

func (i *userInteractor) GetUsers(ctx context.Context) ([]dto.User, error) {
	entities, err := i.UserRepository.GetUsers(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get users", "error", err)
		return nil, err
	}
	users := make([]dto.User, 0)
	for _, entity := range entities {
		tmp := convertUser(entity)
		users = append(users, tmp)
	}

	return users, nil
}

func (i *userInteractor) GetUsersByID(ctx context.Context, id string) (dto.User, error) {
	entity, err := i.UserRepository.GetUsersByID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get user", "id", id, "error", err)
		return dto.User{}, err
	}
	return convertUser(entity), nil
}

func convertUser(u domain.User) dto.User {
	return dto.User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
