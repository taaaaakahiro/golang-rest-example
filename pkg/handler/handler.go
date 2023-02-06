package handler

import (
	handlerTmp "github.com/taaaaakahiro/golang-rest-example/pkg/handler/template"
	v1 "github.com/taaaaakahiro/golang-rest-example/pkg/handler/v1"
	"github.com/taaaaakahiro/golang-rest-example/pkg/handler/version"
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"github.com/taaaaakahiro/golang-rest-example/pkg/service"
	tmpHtml "github.com/taaaaakahiro/golang-rest-example/template"
	"go.uber.org/zap"
)

type Handler struct {
	V1       *v1.Handler
	Version  *version.Handler
	Template *handlerTmp.Handler
}

func NewHandler(
	logger *zap.Logger,
	repo *persistence.Repositories,
	services *service.Service,
	template *tmpHtml.Template,
	ver string,
) *Handler {
	return &Handler{
		V1:       v1.NewHandler(logger, repo, services),
		Version:  version.NewHandler(logger.Named("version"), ver),
		Template: handlerTmp.NewHandler(template),
	}
}
