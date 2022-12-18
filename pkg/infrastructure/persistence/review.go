package persistence

import (
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
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

func (r *ReviewRepository) ListReview() ([]*entity.Review, error) {
	// todo:途中
	return []*entity.Review{}, nil
}
