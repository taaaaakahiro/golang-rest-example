package persistence

import (
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
)

type Repositories struct {
	DB               *io.SQLDatabase
	UserRepository   *UserRepository
	ReviewRepository *ReviewRepository
}

func NewRepositories(db *io.SQLDatabase) (*Repositories, error) {
	return &Repositories{
		DB:               db,
		UserRepository:   NewUserRepository(db),
		ReviewRepository: NewReviewRepository(db),
	}, nil
}
