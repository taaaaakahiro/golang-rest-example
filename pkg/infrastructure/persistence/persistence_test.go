package persistence

import (
	"context"
	"database/sql"
	"fmt"
	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"
	"log"
	"os"
	"testing"

	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
)

const dbname = "example_persistence"

var (
	userRepo       *UserRepository
	reviewRepo     *ReviewRepository
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
	// init
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

	// db
	db, _ := testfixtures.NewTestDatabase(sqlSetting, testDSN)
	// repo
	r, _ := NewRepositories(db)
	userRepo = r.UserRepository
	reviewRepo = r.ReviewRepository

	res := m.Run()
	// after
	err = testfixtures.DropDatabase(dbname, testDB)
	if err != nil {
		log.Fatal(err)
	}

	os.Exit(res)
}
