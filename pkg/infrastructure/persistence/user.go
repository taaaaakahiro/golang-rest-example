package persistence

import (
	"database/sql"

	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/repository"
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
)

type UserRepository struct {
	database *io.SQLDatabase
}

var _ repository.IUserRepository = (*UserRepository)(nil)

func NewUserRepository(db *io.SQLDatabase) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (r *UserRepository) GetUser(userID int) (*entity.User, error) {
	query := `
		SELECT
			id,
			name
		FROM
			users
		WHERE
			id = ?
	`
	stmtOut, err := r.database.Prepare(query)
	if err != nil {
		return nil, err
	}
	var user entity.User
	err = stmtOut.QueryRow(userID).Scan(&user.ID, &user.Name)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, err
		}
	}

	return &user, nil
}
