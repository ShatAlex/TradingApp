package repository

import "github.com/jmoiron/sqlx"

type Authorization interface {
}

type Trade interface {
}

type Type interface {
}

type Repository struct {
	Authorization
	Trade
	Type
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{}
}
