package persistence

import (
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
)

type Repositories struct {
	db     *io.SQLDatabase
	User   *UserRepository
	Review *ReviewRepository
}

func NewRepositories(db *io.SQLDatabase) (*Repositories, error) {
	return &Repositories{
		db:     db,
		User:   NewUserRepository(db),
		Review: NewReviewRepository(db),
	}, nil
}
