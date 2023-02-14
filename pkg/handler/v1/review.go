package v1

import (
	"encoding/json"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/output"
	"go.uber.org/zap"
	"log"
	"net/http"
	"strconv"
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

func (h *Handler) ListReviewsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		v := r.URL.Query()
		if v == nil {
			log.Printf("bad request")
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		page, err := strconv.Atoi(v.Get("page"))
		if err != nil {
			h.logger.Error("failed to convert page", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		perPage, err := strconv.Atoi(v.Get("per_page"))
		if err != nil {
			h.logger.Error("failed to convert per_page", zap.Error(err))
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		reviews, err := h.repo.ReviewRepository.ListReviewsByLimitAndOffset(
			ctx,
			h.repo.DB().Database,
			page,
			perPage,
		)
		if err != nil {
			log.Printf("failed to create review (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(reviews)
		if err != nil {
			log.Printf("failed to marshal (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}
