package server

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/taaaaakahiro/golang-rest-example/pkg/handler"
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"github.com/taaaaakahiro/golang-rest-example/pkg/service"
	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"
	"go.uber.org/zap"
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
)

var (
	testServer     *httptest.Server
	testDB         *sql.DB
	truncateTables = []string{
		"users",
		"reviews",
	}
)

const dbname = "example_server"

var testDSN = fmt.Sprintf(
	"root:password@tcp(localhost:33061)/%s?charset=utf8mb4&parseTime=True",
	dbname,
)

func TestMain(m *testing.M) {
	// before

	cfg, err := config.LoadConfig(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	logger, err := zap.NewDevelopment()
	if err != nil {
		log.Fatal(err)
	}

	initDB, err := sql.Open("mysql", cfg.DB.DSN)
	if err != nil {
		log.Fatal(err)
	}

	err = testfixtures.CreateDatabase(dbname, initDB)
	if err != nil {
		log.Fatal(err)
	}

	testDB, err = sql.Open("mysql", testDSN)
	if err != nil {
		log.Fatal(err)
	}

	err = testfixtures.CreateTables(testDB, testfixtures.Path)
	if err != nil {
		log.Fatal(err)
	}

	sqlSetting := &config.SQLDBSettings{
		SqlDsn:              cfg.DB.DSN,
		SqlMaxOpenConns:     cfg.DB.MaxOpenConns,
		SqlMaxIdleConns:     cfg.DB.MaxIdleConns,
		SqlConnsMaxLifetime: cfg.DB.ConnsMaxLifetime,
	}

	db, err := testfixtures.TestDatabase(sqlSetting, testDSN)
	if err != nil {
		log.Fatal(err)
	}

	repo, _ := persistence.NewRepositories(db)
	services := service.NewService(repo)

	handler := handler.NewHandler(logger, repo, services, "test")
	server := NewServer(handler, &Config{Log: logger}, cfg)
	testServer = httptest.NewServer(server.Router)

	res := m.Run()
	// after

	err = testfixtures.DropDatabase(dbname, testDB)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(res)
}
