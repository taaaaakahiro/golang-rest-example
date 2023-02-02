package persistence

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"

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

func (r *UserRepository) GetUser(ctx context.Context, userID string) (*entity.User, error) {
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
	err = stmtOut.QueryRowContext(ctx, userID).Scan(&user.ID, &user.Name)
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

func (r *UserRepository) ListUsers(ctx context.Context) ([]*entity.User, error) {
	query := `
		SELECT
			id,
			name
		FROM
			users
	`
	stmtOut, err := r.database.Prepare(query)
	if err != nil {
		return nil, err
	}
	rows, err := stmtOut.QueryContext(ctx)
	if err != nil {
		return []*entity.User{}, err
	}

	users := make([]*entity.User, 0)
	for rows.Next() {
		var user entity.User
		err = rows.Scan(&user.ID, &user.Name)
		if err != nil {
			return []*entity.User{}, err
		}
		users = append(users, &user)
	}
	if err = rows.Err(); err != nil {
		return []*entity.User{}, err
	}

	return users, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, name string) (*int, error) {
	query := `
INSERT INTO
	users (name)
VALUE (?)
	`
	stmtOut, err := r.database.Database.Prepare(query)
	if err != nil {
		return nil, err
	}
	result, err := stmtOut.ExecContext(ctx, name)
	if err != nil {
		return nil, err
	}
	insertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	id := int(insertID)

	return &id, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, userID string, name string) error {
	query := `
UPDATE
	users
SET
	name = ?
WHERE
	id = ?
	`
	stmtOut, err := r.database.Database.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmtOut.ExecContext(ctx, name, userID)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) DeleteUser(ctx context.Context, userID string) error {
	query := `
DELETE FROM
	users
WHERE
	id = ?
	`
	stmtOut, err := r.database.Database.Prepare(query)
	if err != nil {
		return err
	}
	result, err := stmtOut.ExecContext(ctx, userID)
	if err != nil {
		return err
	}
	_, err = result.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (r *UserRepository) TxExistUser(ctx context.Context, tx *sql.Tx, userID int) (bool, error) {
	query := `
SELECT EXISTS (
	SELECT
		*
	FROM
	    users
    WHERE
        id = ?
);
`
	stmtOut, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return false, errors.WithStack(err)
	}
	var b bool
	err = stmtOut.QueryRowContext(ctx, userID).Scan(&b)
	if err != nil {
		return false, errors.WithStack(err)
	}

	return b, nil
}
