package repository

import (
	"context"

	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
)

type IUserRepository interface {
	GetUser(ctx context.Context, userID string) (*entity.User, error)
}
