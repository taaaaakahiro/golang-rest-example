package version

import (
	"encoding/json"
	"net/http"

	"go.uber.org/zap"
)

type Handler struct {
	logger *zap.Logger
	res    *response
}

type response struct {
	Version string `json:"version"`
}

func NewHandler(logger *zap.Logger, ver string) *Handler {
	return &Handler{
		logger: logger,
		res: &response{
			Version: ver,
		},
	}
}

func (h *Handler) GetVersion(w http.ResponseWriter, r *http.Request) {
	b, err := json.Marshal(h.res)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		h.logger.Error("failed to marshal version", zap.Error(err))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(b)
}
