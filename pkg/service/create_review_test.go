package service

import (
	"context"
	"github.com/google/go-cmp/cmp"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/input"
	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"
	"testing"
)

func TestReviewService_Create(t *testing.T) {
	users := []struct {
		Id   string
		Name string
	}{
		{
			Id:   "1",
			Name: "user1",
		},
	}

	type args struct {
		inputReview input.Review
	}

	insID := 1
	tests := []struct {
		name    string
		args    args
		want    *int
		wantErr error
	}{
		{
			name: "ok",
			args: args{
				inputReview: input.Review{
					Text:   "text1",
					UserID: 1,
				},
			},
			want:    &insID,
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
		for _, user := range users {
			if err := testfixtures.InsertTable(testDB, "users", interface{}(user)); err != nil {
				t.Errorf("insert error: %s\n", err.Error())
			}
		}
		c := context.Background()
		t.Run(tt.name, func(t *testing.T) {
			got, err := services.ReviewService.Create(c, tt.args.inputReview)
			if diff := cmp.Diff(tt.wantErr, err); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})

	}
}
