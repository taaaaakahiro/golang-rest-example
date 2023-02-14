package service

import (
	"context"
	"github.com/pkg/errors"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"log"
)

type ReviewService struct {
	repo *persistence.Repositories
}

func NewReviewService(r *persistence.Repositories) *ReviewService {
	return &ReviewService{
		repo: r,
	}
}

func (s *ReviewService) Create(ctx context.Context, inputReview input.Review) (*int, error) {
	tx, cancel, err := s.repo.DB().Begin()
	if err != nil {
		return nil, errors.WithStack(err)
	}
	defer cancel()
	defer func() {
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				log.Fatalf("failed to rollback(error:%s)", err)
				return
			}
		} else {
			err = tx.Commit()
			if err != nil {
				log.Fatalf("failed to commit(error:%s)", err)
				return
			}
		}
	}()

	// 存在チェック
	exist, err := s.repo.UserRepository.TxExistUser(ctx, tx, inputReview.UserID)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	if !exist {
		return nil, errors.New("user is not found")
	}

	// レビュー作成
	id, err := s.repo.ReviewRepository.TxCreateReview(ctx, tx, inputReview)
	if err != nil {
		return nil, errors.WithStack(err)
	}

	return id, nil
}
