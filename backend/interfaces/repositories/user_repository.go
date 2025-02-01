package interfaces

import (
	"context"
	"log/slog"
	"strconv"

	"github.com/hashiotoko/go-sample-app/backend/database/sqlc"
	"github.com/hashiotoko/go-sample-app/backend/domain"
	"github.com/hashiotoko/go-sample-app/backend/infrastructure/db"
	repository "github.com/hashiotoko/go-sample-app/backend/usecases/repository_interface"
	"github.com/hashiotoko/go-sample-app/backend/usecases/repository_interface/dto"
)

type UserRepository struct {
	DBClient db.Client
}

func NewUserRepository(client db.Client) repository.UserRepository {
	return &UserRepository{
		DBClient: client,
	}
}

func (r *UserRepository) GetUsers(ctx context.Context) ([]domain.User, error) {
	res, err := r.DBClient.Conn().GetUsers(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get users", "error", err)
		return nil, err
	}
	users := make([]domain.User, 0)
	for _, user := range res {
		tmp := convertUser(user)
		users = append(users, tmp)
	}

	return users, nil
}

func (r *UserRepository) GetUsersByID(ctx context.Context, id string) (domain.User, error) {
	res, err := r.DBClient.Conn().GetUsersByID(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get user", "id", id, "error", err)
		return domain.User{}, err
	}

	return convertUser(res), nil
}

func (r *UserRepository) CreateUser(ctx context.Context, req dto.CreateUserRequest) (domain.User, error) {
	lastIDStr, err := r.DBClient.Conn().GetUserLastID(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get last user id", "error", err)
		return domain.User{}, err
	}
	lastID, err := strconv.Atoi(lastIDStr)
	id := strconv.Itoa(lastID + 1)
	if err != nil {
		slog.ErrorContext(ctx, "failed to convert id", "error", err)
		return domain.User{}, err
	}

	err = r.DBClient.Conn().InsertUser(ctx, sqlc.InsertUserParams{
		ID: id,
		Name: req.Name,
		EmailAddress: req.EmailAddress,
	})
	if err != nil {
		slog.ErrorContext(ctx, "failed to create user", "error", err)
		return domain.User{}, err
	}

	user, err := r.GetUsersByID(ctx, id)
	if err != nil {
		return domain.User{}, err
	}

	return user, nil
}

func convertUser(u sqlc.User) domain.User {
	id, err := strconv.ParseInt(u.ID, 10, 32)
	if err != nil {
		slog.Error("failed to convert id", id, "error", err)
	}
	return domain.User{
		ID:        int32(id),
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
		UpdatedAt: u.UpdatedAt,
	}
}
