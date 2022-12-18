package repository

import "github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"

type IReviewRepository interface {
	ListReview() ([]*entity.Review, error)
}
