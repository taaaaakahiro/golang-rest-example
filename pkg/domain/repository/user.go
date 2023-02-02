package repository

import (
	"context"
	"database/sql"

	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
)

type IUserRepository interface {
	GetUser(ctx context.Context, userID string) (*entity.User, error)
	ListUsers(ctx context.Context) ([]*entity.User, error)
	CreateUser(ctx context.Context, name string) (*int, error)
	UpdateUser(ctx context.Context, userID string, name string) error
	DeleteUser(ctx context.Context, userID string) error
	TxExistUser(ctx context.Context, tx *sql.Tx, userID int) (bool, error)
}
