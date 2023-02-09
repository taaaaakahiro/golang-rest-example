package repository

import (
	"context"
	"database/sql"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"

	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
)

type IReviewRepository interface {
	ListReviews(ctx context.Context, db ContextExecutor, userID int) ([]*entity.Review, error)
	ListReviewsByLimitAndOffset(ctx context.Context, db ContextExecutor, page int, perPage int) ([]*entity.Review, error)
	TxCreateReview(ctx context.Context, tx *sql.Tx, inputReview input.Review) (*int, error)
	GetReview(ctx context.Context, db ContextExecutor, reviewID int) (*entity.Review, error)
}
