package v1

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"

	"go.uber.org/zap"
)

func (h *Handler) GetUserHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.Background()
		vars := mux.Vars(r)
		id := vars["id"]
		user, err := h.repo.User.GetUser(ctx, id)
		if err != nil {
			h.logger.Error("failed to get user", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(user)
		if err != nil {
			h.logger.Error("marshal error", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	})

}
