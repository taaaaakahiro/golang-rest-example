package service

import "github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"

type Service struct {
	ReviewService *ReviewService
}

func NewService(r *persistence.Repositories) *Service {
	return &Service{
		ReviewService: NewReviewService(r),
	}
}
