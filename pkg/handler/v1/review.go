package v1

import (
	"encoding/json"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/output"
	"log"
	"net/http"
)

func (h *Handler) PostReviewHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ctx := req.Context()
		var input input.Review
		decode := json.NewDecoder(req.Body)
		err := decode.Decode(&input)
		if err != nil {
			log.Printf("failed to decode (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id, err := h.services.ReviewService.Create(ctx, input)
		if err != nil {
			log.Printf("failed to create review (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		if id == nil {
			msg := "review is conflict"
			http.Error(w, msg, http.StatusConflict)
			w.WriteHeader(http.StatusConflict)
			return
		}

		output := output.Review{
			ID: *id,
		}
		b, err := json.Marshal(output)
		if err != nil {
			log.Printf("failed to marshal (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte(b))
	}
}
