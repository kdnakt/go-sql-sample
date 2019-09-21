package port

import (
	"context"

	"github.com/kdnakt/go-sql-sample/entity"
)

// UserAccessor is reader and writer for User in datastore
type UserAccessor interface {
	UserReader
	UserWriter
}

// UserReader retrieves User data from datastore
type UserReader interface {
	FindUser(context.Context, int64) (*entity.User, error)
}

// UserWriter stores User data in datastore
type UserWriter interface {
	AddUser(context.Context, *entity.User) error
}
