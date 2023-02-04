package handler

import (
	v1 "github.com/taaaaakahiro/golang-rest-example/pkg/handler/v1"
	"github.com/taaaaakahiro/golang-rest-example/pkg/handler/version"
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"github.com/taaaaakahiro/golang-rest-example/pkg/service"
	"go.uber.org/zap"
)

type Handler struct {
	V1      *v1.Handler
	Version *version.Handler
}

func NewHandler(
	logger *zap.Logger,
	repo *persistence.Repositories,
	services *service.Service,
	ver string,
) *Handler {
	return &Handler{
		V1:      v1.NewHandler(logger, repo, services),
		Version: version.NewHandler(logger.Named("version"), ver),
	}
}
