package persistence

import (
	"context"

	"github.com/pkg/errors"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
	derr "github.com/taaaaakahiro/golang-rest-example/pkg/domain/error"
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

func (r *ReviewRepository) ListReview(ctx context.Context, db repository.ContextExecutor, userID int) ([]*entity.Review, error) {
	// todo:途中
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
