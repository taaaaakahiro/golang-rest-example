package persistence

import (
	"context"
	"github.com/google/go-cmp/cmp/cmpopts"
	derr "github.com/taaaaakahiro/golang-rest-example/pkg/domain/error"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"
)

func TestReviewRepo_ListReviews(t *testing.T) {
	// CleanUp
	if err := testfixtures.TruncateTables(testDB, []string{"reviews"}); err != nil {
		t.Errorf("truncate error: %s\n", err.Error())
	}
	t.Cleanup(func() {
		if err := testfixtures.TruncateTables(testDB, []string{"reviews"}); err != nil {
			t.Errorf("truncate error: %s\n", err.Error())
		}
	})

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
	// Insert Seeds
	for _, review := range reviews {
		if err := testfixtures.InsertTable(testDB, "reviews", interface{}(review)); err != nil {
			t.Errorf("insert error: %s\n", err.Error())
		}
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

		t.Run(tt.name, func(t *testing.T) {
			got, err := reviewRepo.ListReview(context.Background(), testDB, tt.userID)
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
