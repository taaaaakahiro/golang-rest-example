package handler

import (
	"github.com/taaaaakahiro/golang-rest-example/pkg/handler/version"
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"go.uber.org/zap"
)

type Handler struct {
	Version *version.Handler
}

func NewHandler(logger *zap.Logger, repo *persistence.Repositories, ver string) *Handler {
	return &Handler{
		Version: version.NewHandler(logger.Named("version"), ver),
	}
}
