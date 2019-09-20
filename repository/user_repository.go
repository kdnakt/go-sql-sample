package repository

import (
	"context"
	"database/sql"

	"github.com/kdnakt/go-sql-sample/entity"
)

// FindUser gets user from repository
func (repo *Repo) FindUser(ctx context.Context, id int64) (*entity.User, error) {
	u := &entity.User{}
	conn, err := repo.db.Conn(ctx)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	err = conn.QueryRowContext(ctx, `
		SELECT id, name, email, created_at, updated_at FROM user WHERE id = ?
	`, id).Scan(
		&u.ID,
		&u.Name,
		&u.Email,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	switch {
	case err == sql.ErrNoRows:
		return nil, nil
	case err != nil:
		return nil, err
	}
	return u, nil
}

// AddUser adds user to repository
func (repo *Repo) AddUser(ctx context.Context, u *entity.User) error {
	return nil
}
