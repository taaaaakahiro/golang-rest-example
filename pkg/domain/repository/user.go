package repository

import "github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"

type IUserRepository interface {
	GetUser(userID int) (*entity.User, error)
}
