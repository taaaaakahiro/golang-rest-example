package persistence

import (
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
)

type Repositories struct {
	db *io.SQLDatabase
}

func NewRepositories(db *io.SQLDatabase) (*Repositories, error) {
	return &Repositories{
		db: db,
	}, nil
}
