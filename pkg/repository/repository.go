package repository

import (
	trade "github.com/ShatAlex/trading-app"
	"github.com/jmoiron/sqlx"
)

type Authorization interface {
	CreateUser(user trade.User) (int, error)
	GetUserId(username, password_hash string) (int, error)
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
	return &Repository{
		Authorization: NewAuthPostgres(db),
	}
}
