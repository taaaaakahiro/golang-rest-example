package v1

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"

	"go.uber.org/zap"
)

func (h *Handler) GetUserHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id := vars["id"]
		user, err := h.repo.User.GetUser(context.Background(), id)
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

func (h *Handler) ListUsersHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		users, err := h.repo.User.ListUsers(context.Background())
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
	})
}

func (h *Handler) PostUserHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		var user input.User
		decode := json.NewDecoder(req.Body)
		err := decode.Decode(&user)
		if err != nil {
			log.Printf("failed to decode (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		id, err := h.repo.User.CreateUser(context.Background(), user.Name)
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
	})
}

func (h *Handler) DeleteUserHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		vars := mux.Vars(req)
		id := vars["id"]
		err := h.repo.User.DeleteUser(context.Background(), id)
		if err != nil {
			log.Printf("failed to create user (error:%s)", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusNoContent)
	})
}
