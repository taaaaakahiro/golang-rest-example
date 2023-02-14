package server

import (
	"context"
	"github.com/go-chi/chi/v5"
	chiMiddlleware "github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"net"
	"net/http"

	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
	"github.com/taaaaakahiro/golang-rest-example/pkg/handler"
	"github.com/taaaaakahiro/golang-rest-example/pkg/middleware"
	"go.uber.org/zap"
)

type Server struct {
	Router  *chi.Mux
	server  *http.Server
	handler *handler.Handler
	log     *zap.Logger
}

type Config struct {
	Log *zap.Logger
}

func NewServer(registry *handler.Handler, cfg *Config, env *config.Config) *Server {
	s := &Server{
		Router:  chi.NewRouter(),
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

	s.Router.Use(chiMiddlleware.Logger)
	s.Router.Route("/", func(r chi.Router) {
		r.Get("/healthz", s.healthCheckHandler)
		r.Get("/version", s.handler.Version.GetVersion)
		r.Get("/index", s.handler.Template.IndexTemplateHandler)

		// v1
		s.Router.Route("/v1", func(r chi.Router) {
			// 自作の場合
			r.Use(middleware.CORSHeaderMiddleware(env))
			// chi-corsの場合
			r.Use(cors.Handler(cors.Options{
				AllowedOrigins: []string{"*"}, // Use this to allow specific origin hosts
				//AllowedOrigins: []string{"https://*", "http://*"},
				AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
				//AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"}, //CORSHeaderMiddlewareで設定
				AllowCredentials: true,
			}))

			// user
			r.Route("/user", func(r chi.Router) {
				r.Get("/{userID}", s.handler.V1.GetUserHandler())
				r.Get("/all", s.handler.V1.ListUsersHandler())
				r.Post("/", s.handler.V1.PostUserHandler())
				r.Delete("/{userID}", s.handler.V1.DeleteUserHandler())
			})

			// review
			r.Route("/review", func(r chi.Router) {
				r.Post("/", s.handler.V1.PostReviewHandler())
				r.Get("/", s.handler.V1.ListReviewsHandler())
			})

		})

	})

}

func (s *Server) healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
