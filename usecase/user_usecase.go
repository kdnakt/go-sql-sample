package usecase

import (
	"context"

	"github.com/kdnakt/go-sql-sample/entity"
	"github.com/kdnakt/go-sql-sample/usecase/port"
)

// UserCase contains use cases repository to handle entity.User
type UserCase struct {
	ua port.UserAccessor
}

// NewUserCase initialises UserCase
func NewUserCase(ua port.UserAccessor) *UserCase {
	return &UserCase{
		ua: ua,
	}
}

// Save is a use case to save entity.User in a repository
func (au *UserCase) Save(ctx context.Context, name, email string) (int64, error) {
	u := &entity.User{
		Name:  name,
		Email: email,
	}
	err := au.ua.AddUser(ctx, u)
	if err != nil {
		return 0, err
	}
	return u.ID, nil
}

// Find is a use case to retrieve entity.User from a repository
func (au *UserCase) Find(ctx context.Context, id int64) (*entity.User, error) {
	u, err := au.ua.FindUser(ctx, id)
	if err != nil {
		return nil, err
	}
	return u, nil
}
