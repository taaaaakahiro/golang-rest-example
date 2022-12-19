package persistence

import (
	"database/sql"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/taaaaakahiro/golang-rest-example/pkg/domain/entity"
	testfixtures "github.com/taaaaakahiro/golang-rest-example/test_fixtures"
)

const userTable = "users"

func TestUserRepository_GetUser(t *testing.T) {
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
		Id   int
		Name string
	}{
		{Id: 1, Name: "user1"},
		{Id: 2, Name: "user2"},
	}

	for _, user := range users {
		if err := testfixtures.InsertTable(db, "users", interface{}(user)); err != nil {
			t.Errorf("insert error: %s\n", err.Error())
		}
	}

	// TestCase
	tests := []struct {
		name    string
		userID  int
		want    *entity.User
		wantErr error
	}{
		{
			name:   "ok",
			userID: 1,
			want: &entity.User{
				ID: 1, Name: "user1",
			},
			wantErr: nil,
		},
		{
			name:   "ok",
			userID: 2,
			want: &entity.User{
				ID: 2, Name: "user2",
			},
			wantErr: nil,
		},
		{
			name:    "notExistUserId",
			userID:  999,
			want:    nil,
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := userRepo.GetUser(tt.userID)
			if diff := cmp.Diff(tt.wantErr, err); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
			if diff := cmp.Diff(tt.want, got); len(diff) != 0 {
				t.Errorf("mismatch (-want +got):\n%s", diff)
			}
		})
	}

}
