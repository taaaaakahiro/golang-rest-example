package v1

import "github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"

type Handler struct {
	repo *persistence.Repositories
}

func NewHandler(repo *persistence.Repositories) *Handler {
	return &Handler{
		repo: repo,
	}
}
