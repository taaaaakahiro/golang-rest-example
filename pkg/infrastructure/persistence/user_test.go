package persistence

import (
	"database/sql"
	"fmt"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository_GetUser(t *testing.T) {
	db, _ := sql.Open("mysql", mysqlDsn)
	insertFixture(db)

	// fixture := []struct {
	// 	id   int
	// 	name string
	// }{
	// 	{
	// 		id:   1,
	// 		name: "user1",
	// 	},
	// }

	got, err := userRepo.GetUser(1)
	assert.NoError(t, err)
	assert.NotEmpty(t, got)
	fmt.Println(got)

	t.Cleanup(func() {
		stmtOut, err := db.Prepare("set foreign_key_checks = 0;")
		if err != nil {
			log.Fatal(err)
		}

		_, err = stmtOut.Exec()
		if err != nil {
			log.Fatal(err)
		}
		stmtOut, err = db.Prepare("truncate table users")
		if err != nil {
			log.Fatal(err)
		}
		_, err = stmtOut.Exec()
		if err != nil {
			log.Fatal(err)
		}
	})

}

func insertFixture(db *sql.DB) {
	// プリペアードステートメントを使用
	in, err := db.Prepare("INSERT INTO users(name) VALUES(?)")
	if err != nil {
		fmt.Println("データベース接続失敗")
		panic(err.Error())
	} else {
		fmt.Println("データベース接続成功")
	}
	_, err = in.Exec("user1")
	if err != nil {
		panic(err.Error())
	}
}
