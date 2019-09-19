package repository

import (
	"context"
	"reflect"
	"testing"
	"time"

	sqlmock "github.com/DATA-DOG/go-sqlmock"

	"github.com/kdnakt/go-sql-sample/entity"
)

var o = entity.User{
	ID:        1,
	Name:      "kdnakt",
	Email:     "kdnakt@example.com",
	CreatedAt: time.Now(),
	UpdatedAt: time.Now(),
}

func existRows() *sqlmock.Rows {
	return sqlmock.NewRows([]string{
		"id", "name", "email", "created_at", "updated_at",
	}).AddRow(
		o.ID,
		o.Name,
		o.Email,
		o.CreatedAt,
		o.UpdatedAt,
	)
}

func TestRepo_FindUser(t *testing.T) {
	unknownID := int64(999)

	mockdb, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	repo := NewRepo(mockdb)

	mock.ExpectQuery(`
		SELECT id, name, email, created_at, updated_at FROM user WHERE id = ?
	`).WithArgs(o.ID).WillReturnRows(existRows())

	got, err := repo.FindUser(ctx, o.ID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if !reflect.DeepEqual(o, got) {
		t.Errorf("Expected: %v, but got: %v", o, got)
	}

	got2, err := repo.FindUser(ctx, unknownID)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}
	if got2 != nil {
		t.Errorf("Expected: nil, but got: %v", got2)
	}

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("mock has error %v", err)
	}
}