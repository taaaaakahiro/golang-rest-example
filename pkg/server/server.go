package server

import (
	"context"
	"net"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
	"github.com/taaaaakahiro/golang-rest-example/pkg/handler"
	"github.com/taaaaakahiro/golang-rest-example/pkg/middleware"
	"go.uber.org/zap"
)

type Server struct {
	Router  *mux.Router
	server  *http.Server
	handler *handler.Handler
	log     *zap.Logger
}

type Config struct {
	Log *zap.Logger
}

func NewServer(registry *handler.Handler, cfg *Config, env *config.Config) *Server {
	s := &Server{
		Router:  mux.NewRouter(),
		handler: registry,
	}

	if cfg != nil {
		if log := cfg.Log; log != nil {
			s.log = log
		}
	}

	s.registerHandler(env, cfg)
	return s
}

func (s *Server) Serve(listener net.Listener) error {
	server := &http.Server{
		Handler: s.Router,
	}
	s.server = server
	if err := server.Serve(listener); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *Server) GracefulShutdown(ctx context.Context) error {
	return s.server.Shutdown(ctx)
}

func (s *Server) registerHandler(env *config.Config, cnf *Config) {

	// v1
	sr := s.Router.PathPrefix("/v1").Subrouter()
	{
		// user
		sr.Use(mux.CORSMethodMiddleware(s.Router), middleware.CORSHeaderMiddleware(env))
		sr.Handle("/user/{id}", s.handler.V1.GetUserHandler()).Methods(http.MethodGet, http.MethodOptions)
		sr.Handle("/users", s.handler.V1.ListUsersHandler()).Methods(http.MethodGet, http.MethodOptions)
		sr.Handle("/user", s.handler.V1.PostUserHandler()).Methods(http.MethodPost, http.MethodOptions)
		sr.Handle("/user/{id}", s.handler.V1.DeleteUserHandler()).Methods(http.MethodDelete, http.MethodOptions)

		// review
		sr.Handle("/review", s.handler.V1.PostReviewHandler()).Methods(http.MethodPost, http.MethodOptions)
	}

	// common
	s.Router.HandleFunc("/healthz", s.healthCheckHandler).Methods(http.MethodGet)
	s.Router.HandleFunc("/version", s.handler.Version.GetVersion).Methods(http.MethodGet)
	s.Router.HandleFunc("/index", s.handler.Template.IndexTemplateHandler).Methods(http.MethodGet)

}

func (s *Server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
