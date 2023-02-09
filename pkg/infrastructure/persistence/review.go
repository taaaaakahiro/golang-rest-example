package persistence

import (
	"context"
	"database/sql"
	"github.com/pkg/errors"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
	derr "github.com/taaaaakahiro/golang-rest-example/pkg/domain/error"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/repository"
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
)

type ReviewRepository struct {
	database *io.SQLDatabase
}

var _ repository.IReviewRepository = (*ReviewRepository)(nil)

func NewReviewRepository(db *io.SQLDatabase) *ReviewRepository {
	return &ReviewRepository{
		database: db,
	}
}

func (r *ReviewRepository) ListReviews(ctx context.Context, db repository.ContextExecutor, userID int) ([]*entity.Review, error) {
	query := `
SELECT
	id,
	text,
	user_id
FROM
	reviews
WHERE
	user_id = ?
	`

	stmtOut, err := db.PrepareContext(ctx, query)
	if err != nil {
		return []*entity.Review{}, errors.WithStack(err)
	}
	rows, err := stmtOut.QueryContext(ctx, userID)
	if err != nil {
		return []*entity.Review{}, errors.WithStack(err)
	}
	reviews := make([]*entity.Review, 0)
	for rows.Next() {
		var review entity.Review
		err = rows.Scan(
			&review.ID,
			&review.Text,
			&review.UserID,
		)
		if err != nil {
			return []*entity.Review{}, errors.WithStack(err)
		}
		reviews = append(reviews, &review)

	}
	if err = rows.Err(); err != nil {
		return []*entity.Review{}, errors.WithStack(err)
	}
	if len(reviews) == 0 {
		return []*entity.Review{}, derr.ErrReviewNotFound{}
	}

	return reviews, nil
}

func (r *ReviewRepository) TxCreateReview(ctx context.Context, tx *sql.Tx, inputReview input.Review) (*int, error) {
	query := `
INSERT INTO reviews 
    (text, user_id) 
VALUES
    (?,?)
`
	stmtOut, err := tx.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	result, err := stmtOut.ExecContext(ctx, inputReview.Text, inputReview.UserID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	af, err := result.RowsAffected()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if af != 1 {
		return nil, derr.ErrReviewConflict{}
	}
	insID, err := result.LastInsertId()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	id := int(insID)

	return &id, nil
}

func (r *ReviewRepository) GetReview(ctx context.Context, db repository.ContextExecutor, reviewID int) (*entity.Review, error) {
	query := `
SELECT
	id,
	text,
	user_id
FROM
	reviews
WHERE
	id = ?
	`
	stmtOut, err := db.PrepareContext(ctx, query)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer stmtOut.Close()

	var review entity.Review
	err = stmtOut.QueryRowContext(ctx, reviewID).Scan(
		&review.ID,
		&review.Text,
		&review.UserID,
	)
	if err != nil {
		switch err {
		case sql.ErrNoRows:
			return nil, nil
		default:
			return nil, errors.WithStack(err)
		}
	}

	return &review, nil
}

func (r *ReviewRepository) ListReviewsByLimitAndOffset(ctx context.Context, db repository.ContextExecutor, limit int, offset int) ([]*entity.Review, error) {
	query := `
SELECT
    id,
    text,
    user_id
FROM
    reviews
ORDER BY id
LIMIT ?
OFFSET ?
;
`

	stmtOut, err := db.PrepareContext(ctx, query)
	if err != nil {
		return []*entity.Review{}, errors.WithStack(err)
	}

	rows, err := stmtOut.QueryContext(ctx, limit, offset)
	if err != nil {
		return []*entity.Review{}, errors.WithStack(err)
	}
	reviews := make([]*entity.Review, 0)
	for rows.Next() {
		var review entity.Review
		err = rows.Scan(
			&review.ID,
			&review.Text,
			&review.UserID,
		)
		if err != nil {
			return []*entity.Review{}, errors.WithStack(err)
		}
		reviews = append(reviews, &review)
	}

	if err = rows.Err(); err != nil {
		return []*entity.Review{}, errors.WithStack(err)
	}
	if len(reviews) == 0 {
		return []*entity.Review{}, sql.ErrNoRows
	}

	return reviews, nil
}
