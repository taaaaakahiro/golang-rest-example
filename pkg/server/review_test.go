package server

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"
	"io"
	"net/http"
	"testing"
)

func TestServer_ListReviews(t *testing.T) {
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
	reviews := []struct {
		Id     int
		Text   string
		userID int
	}{
		{Id: 1, Text: "text1", userID: 1},
		{Id: 2, Text: "text2", userID: 2},
		{Id: 3, Text: "text3", userID: 3},
	}

	for _, r := range reviews {
		if err := testfixtures.InsertTable(testDB, "reviews", interface{}(r)); err != nil {
			t.Errorf("insert error: %s\n", err.Error())
		}
	}

	t.Run("ListReviews", func(t *testing.T) {
		t.Run("ok: page 1 & perPage 1", func(t *testing.T) {
			page := 1
			perPage := 0
			ep := fmt.Sprintf("%s/v1/review", testServer.URL)
			req, err := http.NewRequest("GET", ep, nil)
			assert.NoError(t, err)
			req.Header.Set("Content-Type", "application/json")

			q := req.URL.Query()
			q.Add("page", fmt.Sprint(page))
			q.Add("per_page", fmt.Sprint(perPage))
			req.URL.RawQuery = q.Encode()

			client := &http.Client{}
			res, err := client.Do(req)
			assert.NoError(t, err)
			assert.NotEmpty(t, res)

			body, err := io.ReadAll(res.Body)
			assert.NoError(t, err)
			assert.NotEmpty(t, body)
			defer res.Body.Close()

			var data []*entity.Review
			err = json.Unmarshal(body, &data)
			assert.NoError(t, err)
			assert.Equal(t, 1, data[0].ID)
			assert.Equal(t, "text1", data[0].Text)
			assert.Equal(t, 1, data[0].UserID)
		})
	})
}
