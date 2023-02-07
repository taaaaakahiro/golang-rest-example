package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"
	"log"
	"os"
	"testing"
)

var (
	services       *Service
	testDB         *sql.DB
	truncateTables = []string{
		"users",
		"reviews",
	}
)

const dbname = "example"

func TestMain(m *testing.M) {
	// before
	cfg, _ := config.LoadConfig(context.Background())
	sqlSetting := &config.SQLDBSettings{
		SqlDsn:              cfg.DB.DSN,
		SqlMaxOpenConns:     cfg.DB.MaxOpenConns,
		SqlMaxIdleConns:     cfg.DB.MaxIdleConns,
		SqlConnsMaxLifetime: cfg.DB.ConnsMaxLifetime,
	}
	mysqlDsn := fmt.Sprintf(
		"root:password@tcp(localhost:33061)/%s?charset=utf8&parseTime=true",
		dbname,
	)
	testDB, _ = sql.Open("mysql", mysqlDsn)
	if err := testDB.Ping(); err != nil {
		log.Fatal(err)
	}
	err := testfixtures.CreateTables(testDB, testfixtures.Path)
	if err != nil {
		log.Fatal(err)
	}

	db, _ := io.NewDatabase(sqlSetting)
	r, _ := persistence.NewRepositories(db)
	services = NewService(r)

	res := m.Run()
	// after

	os.Exit(res)

}
