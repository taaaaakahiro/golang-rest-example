package repository

import (
	"context"

	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
)

type IUserRepository interface {
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	CreateUser(ctx context.Context, name string) (*int, error)
	UpdateUser(ctx context.Context, userID string, name string) error
	DeleteUser(ctx context.Context, userID string) error
}
