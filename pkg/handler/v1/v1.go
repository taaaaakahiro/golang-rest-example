package v1

import (
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	repo   *persistence.Repositories
}

func NewHandler(logger *zap.Logger, repo *persistence.Repositories) *Handler {
	return &Handler{
		repo: repo,
	}
}
