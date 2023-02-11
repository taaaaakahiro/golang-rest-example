package v1

import (
	"context"
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"

	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"

	"go.uber.org/zap"
)

func (h *Handler) GetUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userID := chi.URLParam(r, "userID")
		user, err := h.repo.UserRepository.GetUser(context.Background(), userID)
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
	}
}

func (h *Handler) ListUsersHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := r.Context()
		users, err := h.repo.UserRepository.ListUsers(c)
		if err != nil {
			h.logger.Error("failed to get user", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		b, err := json.Marshal(users)
		if err != nil {
			h.logger.Error("marshal error", zap.Error(err))
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		w.Write(b)
	}
}

func (h *Handler) PostUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		var user input.User
		decode := json.NewDecoder(req.Body)
		err := decode.Decode(&user)
		if err != nil {
			log.Printf("failed to decode (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id, err := h.repo.UserRepository.CreateUser(context.Background(), user.Name)
		if err != nil {
			log.Printf("failed to create user (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		if *id == 0 {
			log.Printf("user is conflict")
			w.WriteHeader(http.StatusConflict)
			return
		}

		byte, err := json.Marshal(user)
		if err != nil {
			log.Printf("failed to marshal user (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write(byte)
	}
}

func (h *Handler) DeleteUserHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		userID := chi.URLParam(req, "userID")
		err := h.repo.UserRepository.DeleteUser(context.Background(), userID)
		if err != nil {
			log.Printf("failed to create user (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	}
}
