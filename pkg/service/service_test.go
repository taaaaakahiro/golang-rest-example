package service

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
	"github.com/taaaaakahiro/golang-rest-example/pkg/infrastructure/persistence"
	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"
	"log"
	"os"
	"testing"
)

const dbname = "example_service"

var (
	services       *Service
	testDB         *sql.DB
	truncateTables = []string{
		"users",
		"reviews",
	}
	testDSN = fmt.Sprintf(
		"root:password@tcp(localhost:33061)/%s?charset=utf8mb4&parseTime=True",
		dbname,
	)
)

func TestMain(m *testing.M) {
	// before
	cfg, _ := config.LoadConfig(context.Background())
	sqlSetting := &config.SQLDBSettings{
		SqlDsn:              cfg.DB.DSN,
		SqlMaxOpenConns:     cfg.DB.MaxOpenConns,
		SqlMaxIdleConns:     cfg.DB.MaxIdleConns,
		SqlConnsMaxLifetime: cfg.DB.ConnsMaxLifetime,
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

	db, _ := testfixtures.NewTestDatabase(sqlSetting, testDSN)
	r, _ := persistence.NewRepositories(db)
	services = NewService(r)

	res := m.Run()
	// after
	err = testfixtures.DropDatabase(dbname, testDB)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(res)

}
