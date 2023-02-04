package v1

import (
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"github.com/taaaaakahiro/golang-rest-example/pkg/service"
	"go.uber.org/zap"
)

type Handler struct {
	logger   *zap.Logger
	repo     *persistence.Repositories
	services *service.Service
}

func NewHandler(logger *zap.Logger, repo *persistence.Repositories, services *service.Service) *Handler {
	return &Handler{
		logger:   logger,
		repo:     repo,
		services: services,
	}
}
