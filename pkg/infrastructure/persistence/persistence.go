package persistence

import (
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
)

type Repositories struct {
	db               *io.SQLDatabase
	UserRepository   *UserRepository
	ReviewRepository *ReviewRepository
}

func NewRepositories(db *io.SQLDatabase) (*Repositories, error) {
	return &Repositories{
		db:               db,
		UserRepository:   NewUserRepository(db),
		ReviewRepository: NewReviewRepository(db),
	}, nil
}

func (r *Repositories) DB() *io.SQLDatabase {
	return r.db
}
