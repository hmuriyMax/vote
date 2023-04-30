package db

import (
	"github.com/hmuriyMax/vote/internal/repo"
	"github.com/jmoiron/sqlx"
)

type Postgres struct {
	db *sqlx.DB
}

func NewPostgres(db *sqlx.DB) repo.Repository {
	return &Postgres{
		db: db,
	}
}
