package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/taaaaakahiro/golang-rest-example/pkg/service"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
	"github.com/taaaaakahiro/golang-rest-example/pkg/handler"
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
	"go.uber.org/zap"
)

var (
	testServer *httptest.Server
	testDB     *sql.DB
)

const dbname = "example"

func TestMain(m *testing.M) {
	// before
	dsn := fmt.Sprintf(
		"root:password@tcp(localhost:33061)/%s?charset=utf8&parseTime=true",
		dbname,
	)
	testDB, _ = sql.Open("mysql", dsn)

	logger, _ := zap.NewDevelopment()
	cfg, _ := config.LoadConfig(context.Background())
	sqlSetting := &config.SQLDBSettings{
		SqlDsn:              cfg.DB.DSN,
		SqlMaxOpenConns:     cfg.DB.MaxOpenConns,
		SqlMaxIdleConns:     cfg.DB.MaxIdleConns,
		SqlConnsMaxLifetime: cfg.DB.ConnsMaxLifetime,
	}
	db, _ := io.NewDatabase(sqlSetting)
	repo, _ := persistence.NewRepositories(db)
	services := service.NewService(repo)

	handler := handler.NewHandler(logger, repo, services, "test")
	server := NewServer(handler, &Config{Log: logger}, cfg)
	testServer = httptest.NewServer(server.Router)

	res := m.Run()
	// after

	os.Exit(res)
}
