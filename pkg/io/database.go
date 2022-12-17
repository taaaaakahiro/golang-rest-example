package io

import (
	"context"
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
	errs "github.com/pkg/errors"
)

const (
	queryTimeoutSec = 3
)

type MySQLSettings interface {
	DSN() string
	MaxOpenConns() int
	MaxIdleConns() int
	ConnsMaxLifetime() int
}

type SQLDatabase struct {
	Database *sql.DB
}

func NewDatabase(setting MySQLSettings) (*SQLDatabase, error) {
	db, err := sql.Open("mysql", setting.DSN())
	if err != nil {
		return nil, errs.WithStack(err)
	}

	// check config
	if setting.MaxOpenConns() <= 0 {
		return nil, errs.WithStack(errs.New("require set max open conns"))
	}
	if setting.MaxIdleConns() <= 0 {
		return nil, errs.WithStack(errs.New("require set max idle conns"))
	}
	if setting.ConnsMaxLifetime() <= 0 {
		return nil, errs.WithStack(errs.New("require set conns max lifetime"))
	}
	db.SetMaxOpenConns(setting.MaxOpenConns())
	db.SetMaxIdleConns(setting.MaxIdleConns())
	db.SetConnMaxLifetime(time.Duration(setting.ConnsMaxLifetime()) * time.Second)

	return &SQLDatabase{Database: db}, nil
}

func (d *SQLDatabase) Ping() error {
	return d.Database.Ping()
}

func (d *SQLDatabase) Begin() (*sql.Tx, context.CancelFunc, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	tx, err := d.Database.BeginTx(ctx, &sql.TxOptions{
		Isolation: 0,
		ReadOnly:  false,
	})

	return tx, cancel, err
}

func (d *SQLDatabase) Close() error {
	return d.Database.Close()
}

func (d *SQLDatabase) Prepare(query string) (*sql.Stmt, error) {
	if d.Database == nil {
		return nil, errDoesNotDB()
	}

	ctx, cancel := context.WithTimeout(context.Background(), queryTimeoutSec*time.Second)
	defer cancel()
	stmt, err := d.Database.PrepareContext(ctx, query)

	return stmt, err
}

func (d *SQLDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	if d.Database == nil {
		return nil, errDoesNotDB()
	}

	ctx, cancel := context.WithTimeout(context.Background(), queryTimeoutSec*time.Second)
	defer cancel()
	res, err := d.Database.ExecContext(ctx, query, args...)

	return res, err
}

func errDoesNotDB() error {
	return errs.New("database does not exist. Please Open() first")
}
