package handler

import (
	v1 "github.com/taaaaakahiro/golang-rest-example/pkg/handler/v1"
	"github.com/taaaaakahiro/golang-rest-example/pkg/handler/version"
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"go.uber.org/zap"
)

type Handler struct {
	V1      *v1.Handler
	Version *version.Handler
}

func NewHandler(logger *zap.Logger, repo *persistence.Repositories, ver string) *Handler {
	return &Handler{
		V1:      v1.NewHandler(repo),
		Version: version.NewHandler(logger.Named("version"), ver),
	}
}
