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

type TypeTrade interface {
	Create(userId int, typeTrade trade.TypeTrade) (int, error)
	GetAll(userId int) ([]trade.TypeTrade, error)
	GetTypeById(userId, typeId int) (trade.TypeTrade, error)
	Delete(userId, typeId int) error
	Update(userId, typeId int, typeTrade trade.TypeTrade) error
}

type Repository struct {
	Authorization
	Trade
	TypeTrade
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TypeTrade:     NewTypeTradePostgres(db),
	}
}
