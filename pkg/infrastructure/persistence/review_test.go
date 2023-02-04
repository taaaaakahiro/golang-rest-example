package persistence

import (
	"context"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/stretchr/testify/assert"
	derr "github.com/taaaaakahiro/golang-rest-example/pkg/domain/error"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"
)

func TestReviewRepo_ListReviews(t *testing.T) {
	// Seeds
	reviews := []struct {
		id      int
		text    string
		user_id int
	}{
		{
			id:      1,
			text:    "text1",
			user_id: 1,
		},
		{
			id:      2,
			text:    "text2",
			user_id: 1,
		},
		{
			id:      3,
			text:    "text3",
			user_id: 2,
		},
	}

	// TestCase
	tests := []struct {
		name    string
		userID  int
		want    []*entity.Review
		wantErr error
	}{
		{
			name:   "ok: userId 1",
			userID: 1,
			want: []*entity.Review{
				{ID: 1, Text: "text1", UserID: 1},
				{ID: 2, Text: "text2", UserID: 1},
			},
			wantErr: nil,
		},
		{
			name:   "ok: userId 2",
			userID: 2,
			want: []*entity.Review{
				{ID: 3, Text: "text3", UserID: 2},
			},
			wantErr: nil,
		},
		{
			name:    "not exist userId",
			userID:  999,
			want:    []*entity.Review{},
			wantErr: derr.ErrReviewNotFound{},
		},
	}

	for _, tt := range tests {
		// CleanUp
		if err := testfixtures.TruncateTables(testDB, truncateTables); err != nil {
			t.Errorf("truncate error: %s\n", err.Error())
		}
		t.Cleanup(func() {
			if err := testfixtures.TruncateTables(testDB, truncateTables); err != nil {
				t.Errorf("truncate error: %s\n", err.Error())
			}
		})
		// Insert Seeds
		for _, review := range reviews {
			if err := testfixtures.InsertTable(testDB, "reviews", interface{}(review)); err != nil {
				t.Errorf("insert error: %s\n", err.Error())
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := reviewRepo.ListReviews(context.Background(), testDB, tt.userID)
			opt := cmpopts.EquateErrors()
			if diff := cmp.Diff(tt.wantErr, err, opt); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}

}

func TestReviewRepository_TxCreateReview(t *testing.T) {
	users := []struct {
		Id   string
		Name string
	}{
		{Id: "1", Name: "user1"},
		{Id: "2", Name: "user2"},
	}

	insId1 := 1

	tests := []struct {
		name        string
		inputReview input.Review
		want        *int
		wantErr     error
	}{
		{
			name: "ok",
			inputReview: input.Review{
				Text:   "text1",
				UserID: 1,
			},
			want:    &insId1,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		if err := testfixtures.TruncateTables(testDB, truncateTables); err != nil {
			t.Errorf("truncate error: %s\n", err.Error())
		}
		t.Cleanup(func() {
			if err := testfixtures.TruncateTables(testDB, truncateTables); err != nil {
				t.Errorf("truncate error: %s\n", err.Error())
			}
		})
		for _, user := range users {
			if err := testfixtures.InsertTable(testDB, "users", interface{}(user)); err != nil {
				t.Errorf("insert error: %s\n", err.Error())
			}
		}

		c := context.Background()

		t.Run(tt.name, func(t *testing.T) {
			tx, err := testDB.BeginTx(c, nil)
			assert.NoError(t, err)
			defer func() {
				if err != nil {
					_ = tx.Rollback()
				} else {
					_ = tx.Commit()
				}
			}()

			got, err := reviewRepo.TxCreateReview(c, tx, tt.inputReview)
			opt := cmpopts.EquateErrors()
			if diff := cmp.Diff(tt.wantErr, err, opt); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}

		})
	}
}

func TestReviewRepository_GetReview(t *testing.T) {
	// Seeds
	reviews := []struct {
		id      int
		text    string
		user_id int
	}{
		{
			id:      1,
			text:    "text1",
			user_id: 1,
		},
		{
			id:      2,
			text:    "text2",
			user_id: 1,
		},
		{
			id:      3,
			text:    "text3",
			user_id: 2,
		},
	}

	// TestCase
	tests := []struct {
		name    string
		userID  int
		want    *entity.Review
		wantErr error
	}{
		{
			name:   "ok: userId 1",
			userID: 1,
			want: &entity.Review{
				ID:     1,
				Text:   "text1",
				UserID: 1,
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		// CleanUp
		if err := testfixtures.TruncateTables(testDB, truncateTables); err != nil {
			t.Errorf("truncate error: %s\n", err.Error())
		}
		t.Cleanup(func() {
			if err := testfixtures.TruncateTables(testDB, truncateTables); err != nil {
				t.Errorf("truncate error: %s\n", err.Error())
			}
		})
		// Insert Seeds
		for _, review := range reviews {
			if err := testfixtures.InsertTable(testDB, "reviews", interface{}(review)); err != nil {
				t.Errorf("insert error: %s\n", err.Error())
			}
		}

		t.Run(tt.name, func(t *testing.T) {
			got, err := reviewRepo.GetReview(context.Background(), testDB, tt.userID)
			opt := cmpopts.EquateErrors()
			if diff := cmp.Diff(tt.wantErr, err, opt); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}
}
