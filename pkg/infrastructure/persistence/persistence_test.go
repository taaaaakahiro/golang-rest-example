package persistence

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
)

const dbname = "example"

var (
	userRepo       *UserRepository
	reviewRepo     *ReviewRepository
	testDB         *sql.DB
	truncateTables = []string{
		"users",
		"reviews",
	}
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
	mysqlDsn := fmt.Sprintf(
		"root:password@tcp(localhost:33061)/%s?charset=utf8&parseTime=true",
		dbname,
	)
	testDB, _ = sql.Open("mysql", mysqlDsn)
	if err := testDB.Ping(); err != nil {
		log.Fatal(err)
	}

	// db
	db, _ := io.NewDatabase(sqlSetting)
	// repo
	r, _ := NewRepositories(db)
	userRepo = r.UserRepository
	reviewRepo = r.ReviewRepository

	res := m.Run()
	// after

	os.Exit(res)
}
