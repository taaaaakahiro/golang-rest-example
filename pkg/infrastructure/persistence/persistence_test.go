package persistence

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/taaaaakahiro/golang-rest-example/pkg/config"
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
)

const dbname = "example"

var (
	userRepo *UserRepository
	mysqlDsn string
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
	mysqlDsn = fmt.Sprintf(
		"root:password@tcp(localhost:33061)/%s?charset=utf8&parseTime=true",
		dbname,
	)
	// db
	db, _ := io.NewDatabase(sqlSetting)
	// repo
	r, _ := NewRepositories(db)
	userRepo = r.User

	res := m.Run()
	// after

	os.Exit(res)
}
