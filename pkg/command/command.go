package command

import (
	"context"
	"fmt"
	"github.com/taaaaakahiro/golang-rest-example/pkg/service"
	"github.com/taaaaakahiro/golang-rest-example/template"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
	"github.com/taaaaakahiro/golang-rest-example/pkg/handler"
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
	"github.com/taaaaakahiro/golang-rest-example/pkg/server"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"honnef.co/go/tools/lintcmd/version"
)

const (
	exitOK  = 0
	exitErr = 1
)

func Run() {
	os.Exit(run(context.Background()))
}

func run(ctx context.Context) int {
	// Logger
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to setup logger: %s\n", err)
		return exitErr
	}
	defer logger.Sync()
	logger = logger.With(zap.String("version", version.Version))

	// Config
	cfg, err := config.LoadConfig(ctx)
	if err != nil {
		logger.Error("failed to load config", zap.Error(err))
		return exitErr
	}

	// init listener
	listener, err := net.Listen("tcp", cfg.Address())
	if err != nil {
		logger.Error("failed to listen port", zap.Int("port", cfg.Server.Port), zap.Error(err))
		return exitErr
	}
	logger.Info("server start listening", zap.Int("port", cfg.Server.Port))

	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	// DB
	sqlSetting := &config.SQLDBSettings{
		SqlDsn:              cfg.DB.DSN,
		SqlMaxOpenConns:     cfg.DB.MaxOpenConns,
		SqlMaxIdleConns:     cfg.DB.MaxIdleConns,
		SqlConnsMaxLifetime: cfg.DB.ConnsMaxLifetime,
	}
	db, err := io.NewDatabase(sqlSetting)
	if err != nil {
		logger.Error("failed to connect db", zap.Error(err))
		return exitErr
	} else {
		logger.Info("successed to connect db")
	}
	if err = db.Ping(); err != nil {
		logger.Error("failed to ping mysql db", zap.Error(err))
		return exitErr
	}

	// Repository
	repositories, err := persistence.NewRepositories(db)
	if err != nil {
		logger.Error("failed to create repositories", zap.Error(err))
		return exitErr
	}

	// Services
	services := service.NewService(repositories)

	// Template
	templates := template.NewTemplate()

	registry := handler.NewHandler(logger, repositories, services, templates, version.Version)
	httpServer := server.NewServer(registry, &server.Config{Log: logger}, cfg)
	wg, ctx := errgroup.WithContext(ctx)
	wg.Go(func() error {
		return httpServer.Serve(listener)
	})

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGTERM, os.Interrupt)
	select {
	case <-sigCh:
	case <-ctx.Done():
	}

	if err := httpServer.GracefulShutdown(ctx); err != nil {
		return exitErr
	}

	cancel()
	if err := wg.Wait(); err != nil {
		return exitErr
	}

	return exitOK
}
