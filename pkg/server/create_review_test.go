package server

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"
	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"
	"io"
	"net/http"
	"testing"
)

func TestServer_CreateReview(t *testing.T) {
	// CleanUp
	if err := testfixtures.TruncateTables(testDB, truncateTables); err != nil {
		t.Errorf("truncate error: %s\n", err.Error())
	}
	t.Cleanup(func() {
		if err := testfixtures.TruncateTables(testDB, truncateTables); err != nil {
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
		if err := testfixtures.InsertTable(testDB, "users", interface{}(user)); err != nil {
			t.Errorf("insert error: %s\n", err.Error())
		}
	}

	t.Run("CreateReview", func(t *testing.T) {
		t.Run("ok", func(t *testing.T) {
			r := input.Review{
				Text:   "text1",
				UserID: 1,
			}
			b, _ := json.Marshal(r)

			req, err := http.NewRequest("POST", testServer.URL+`/v1/review`, bytes.NewReader(b))
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
			assert.JSONEq(t, `
{
	"id": 1
}
`, string(body))

		})
	})
}
