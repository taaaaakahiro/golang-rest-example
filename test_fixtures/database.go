package testfixtures

import (
	"bytes"
	"database/sql"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"

	errs "github.com/pkg/errors"
	"github.com/taaaaakahiro/golang-rest-example/pkg/io"
)

var (
	testDBs       = make(map[string]*io.SQLDatabase)
	_, file, _, _ = runtime.Caller(0)
	Path          = filepath.Join(filepath.Dir(file), "/1_ddl.sql")
)

func CreateDatabase(dbName string, testDB *sql.DB) error {
	query := fmt.Sprintf("CREATE SCHEMA IF NOT EXISTS %s DEFAULT CHARACTER SET utf8mb4", dbName)
	_, err := testDB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func DropDatabase(dbName string, testDB *sql.DB) error {
	query := fmt.Sprintf("DROP DATABASE IF EXISTS %s", dbName)
	_, err := testDB.Exec(query)
	if err != nil {
		return err
	}
	return nil
}

func TruncateTables(db *sql.DB, tableNames []string) error {
	_, err := db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		return err
	}

	for _, tableName := range tableNames {
		query := fmt.Sprintf("TRUNCATE TABLE %s", tableName)
		_, err = db.Exec(query)
		if err != nil {
			return err
		}
	}
	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		return err
	}
	return nil
}

func InsertTable(db *sql.DB, tableName string, fixture interface{}) error {
	_, err := db.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		return err
	}

	reflectValue := reflect.ValueOf(fixture)
	reflectType := reflect.TypeOf(fixture)

	var (
		keys   []string
		values []interface{}
	)
	for i := 0; i < reflectValue.NumField(); i++ {
		keys = append(keys, camelToSnake(reflectType.Field(i).Name))
		values = append(values, getFieldValue(reflectValue.Field(i)))
	}

	query := "INSERT INTO " + tableName + "(" + strings.Join(keys, ",") + ") VALUES (?" + strings.Repeat(",?", len(keys)-1) + ");"
	_, err = db.Exec(query, values...)
	if err != nil {
		return err
	}

	_, err = db.Exec("SET FOREIGN_KEY_CHECKS = 1")
	if err != nil {
		return err
	}
	return nil
}

func getFieldValue(v reflect.Value) interface{} {
	kind := v.Kind()
	if kind == reflect.String {
		return v.String()
	} else if kind == reflect.Int || kind == reflect.Int32 || kind == reflect.Int64 {
		return v.Int()
	} else if kind == reflect.Float32 || kind == reflect.Float64 {
		return v.Float()
	} else if kind == reflect.Bool {
		return v.Bool()
	} else if kind == reflect.Ptr {
		if v.IsNil() {
			return nil
		} else {
			return getFieldValue(v.Elem())
		}
	} else {
		return v.String()
	}

}

func CreateTables(testDB *sql.DB, path string) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	ddl := bytes.NewBuffer(content).String()
	ls := strings.Split(string(ddl), ";")
	for _, item := range ls {
		_, err := testDB.Exec(item)
		if err != nil {
			fmt.Println(err)
		}

	}
	_, err = testDB.Exec(ddl)
	if err != nil {
		log.Println(err.Error())
	}
	return nil
}

func NewTestDatabase(setting io.MySQLSettings, testDSN string) (*io.SQLDatabase, error) {
	db, err := sql.Open("mysql", testDSN)
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

	return &io.SQLDatabase{Database: db}, nil
}

func camelToSnake(s string) string {
	if s == "" {
		return s
	}

	delimiter := "_"
	sLen := len(s)
	var snake string
	for i, current := range s {
		if i > 0 && i+1 < sLen {
			if current >= 'A' && current <= 'Z' {
				next := s[i+1]
				prev := s[i-1]
				if (next >= 'a' && next <= 'z') || (prev >= 'a' && prev <= 'z') {
					snake += delimiter
				}
			}
		}
		snake += string(current)
	}

	snake = strings.ToLower(snake)
	return snake
}

func TestDatabase(setting io.MySQLSettings, testDSN string) (*io.SQLDatabase, error) {
	if db, ok := testDBs[testDSN]; ok {
		return db, nil
	}

	db, err := sql.Open("mysql", testDSN)
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

	testDB := &io.SQLDatabase{Database: db}
	testDBs[testDSN] = testDB

	return testDB, nil
}
