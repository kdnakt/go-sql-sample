// +build e2e

package e2e

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/kdnakt/go-sql-sample/repository"
	"github.com/kdnakt/go-sql-sample/usecase"
)

func TestUserCase_Save(t *testing.T) {
	okName := "kdnakt"
	okEmail := "kdnakt@example.com"
	type args struct {
		name, email string
	}
	okArgs := args{
		name:  okName,
		email: okEmail,
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "correct",
			args: okArgs,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := os.Getenv("MYSQL_USER")
			p := os.Getenv("MYSQL_PORT")
			db, err := sql.Open(
				"mysql",
				fmt.Sprintf("%s:@(localhost:%s)/sql_sample?parseTime=true&loc=Asia%%2FTokyo", u, p),
			)
			if err != nil {
				log.Fatal(err)
			}
			ctx := context.Background()
			repo := repository.NewRepo(db)
			uc := usecase.NewUserCase(repo)

			got, err := uc.Save(ctx, tt.args.name, tt.args.email)
			if err != nil {
				t.Errorf("want no err, but has error %#v", err)
			}
			if got == 0 {
				t.Error("ID was 0")
			}
		})
	}
}
