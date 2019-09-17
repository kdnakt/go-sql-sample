package repository

import "database/sql"

// Repo is a simple wrapper of sql.DB
type Repo struct {
	db *sql.DB
}

// NewRepo is to create *Repo object
func NewRepo(db *sql.DB) *Repo {
	return &Repo{db}
}
