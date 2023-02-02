package server

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"

	"github.com/stretchr/testify/assert"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"
)

const userTable = "users"

func TestServer_GetUser(t *testing.T) {
	db, _ := sql.Open("mysql", mysqlDsn)

	// CleanUp
	if err := testfixtures.TruncateTables(db, []string{userTable}); err != nil {
		t.Errorf("truncate error: %s\n", err.Error())
	}
	t.Cleanup(func() {
		if err := testfixtures.TruncateTables(db, []string{userTable}); err != nil {
			t.Errorf("truncate error: %s\n", err.Error())
		}
	})

	// Fixture
	users := []struct {
		Id   string
		Name string
	}{
		{Id: "1", Name: "user1"},
		{Id: "2", Name: "user2"},
	}

	for _, user := range users {
		if err := testfixtures.InsertTable(db, "users", interface{}(user)); err != nil {
			t.Errorf("insert error: %s\n", err.Error())
		}
	}

	t.Run("GetUser", func(t *testing.T) {
		type inputUser struct {
			id string
		}

		t.Run("ok", func(t *testing.T) {
			r := inputUser{
				id: "1",
			}
			b, _ := json.Marshal(r)

			req, err := http.NewRequest("GET", testServer.URL+`/v1/user/1`, bytes.NewReader(b))
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			res, err := client.Do(req)
			assert.NoError(t, err)
			assert.NotEmpty(t, res)

			body, err := io.ReadAll(res.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)
			defer res.Body.Close()

			var data entity.User
			_ = json.Unmarshal(body, &data)
			assert.Equal(t, "1", data.ID)
			assert.Equal(t, "user1", data.Name)
		})
	})
}

func TestServer_CreateUser(t *testing.T) {
	db, _ := sql.Open("mysql", mysqlDsn)

	// CleanUp
	if err := testfixtures.TruncateTables(db, []string{userTable}); err != nil {
		t.Errorf("truncate error: %s\n", err.Error())
	}
	t.Cleanup(func() {
		if err := testfixtures.TruncateTables(db, []string{userTable}); err != nil {
			t.Errorf("truncate error: %s\n", err.Error())
		}
	})

	t.Run("CreateUser", func(t *testing.T) {
		r := input.User{
			Name: "user",
		}
		b, _ := json.Marshal(r)

		req, err := http.NewRequest("POST", testServer.URL+`/v1/user`, bytes.NewReader(b))
		assert.NoError(t, err)
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		res, err := client.Do(req)
		assert.NoError(t, err)
		assert.NotEmpty(t, res)
		assert.Equal(t, http.StatusOK, res.StatusCode)

		body, err := io.ReadAll(res.Body)
		assert.NoError(t, err)
		assert.NotEmpty(t, body)
		defer res.Body.Close()

		var data input.User
		_ = json.Unmarshal(body, &data)
		assert.Equal(t, "user", data.Name)
	})

}

func TestServer_DeleteUser(t *testing.T) {
	db, _ := sql.Open("mysql", mysqlDsn)

	// CleanUp
	if err := testfixtures.TruncateTables(db, []string{userTable}); err != nil {
		t.Errorf("truncate error: %s\n", err.Error())
	}
	t.Cleanup(func() {
		if err := testfixtures.TruncateTables(db, []string{userTable}); err != nil {
			t.Errorf("truncate error: %s\n", err.Error())
		}
	})

	// Fixture
	users := []struct {
		Id   string
		Name string
	}{
		{Id: "1", Name: "user1"},
		{Id: "2", Name: "user2"},
	}

	for _, user := range users {
		if err := testfixtures.InsertTable(db, "users", interface{}(user)); err != nil {
			t.Errorf("insert error: %s\n", err.Error())
		}
	}

	t.Run("DeleteUser", func(t *testing.T) {
		t.Run("ok: user 1", func(t *testing.T) {
			id := 1
			req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/v1/user/%d", testServer.URL, id), nil)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{}
			res, err := client.Do(req)
			assert.NoError(t, err)
			assert.NotEmpty(t, res)
			assert.Equal(t, http.StatusNoContent, res.StatusCode)
		})
	})
}
