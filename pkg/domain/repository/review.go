package repository

import (
	"context"
	"database/sql"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"

	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
)

type IReviewRepository interface {
	ListReview(ctx context.Context, db ContextExecutor, userID int) ([]*entity.Review, error)
	TxCreateReview(ctx context.Context, tx *sql.Tx, inputReview input.Review) (*int, error)
}
