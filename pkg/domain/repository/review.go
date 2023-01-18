package repository

import (
	"context"

	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
)

type IReviewRepository interface {
	ListReview(ctx context.Context, db ContextExecutor, userID int) ([]*entity.Review, error)
}
