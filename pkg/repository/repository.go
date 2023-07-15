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
	Create(userId int, trade trade.Trade) (int, error)
	GetAll(userId int) ([]trade.Trade, error)
	GetTradeById(userId, tradeId int) (trade.Trade, error)
	Delete(userId, tradeId int) error
	Update(userId, tradeId int, trade trade.UpdateTradeInput) error
}

type TypeTrade interface {
	Create(userId int, typeTrade trade.TypeTrade) (int, error)
	GetAll(userId int) ([]trade.TypeTrade, error)
	GetTypeById(userId, typeId int) (trade.TypeTrade, error)
	Delete(userId, typeId int) error
	Update(userId, typeId int, typeTrade trade.TypeTrade) error
}

type Portfolio interface {
	BuyTicker(userId int, input trade.BuySellTickerInput, price float64) (int, error)
	SellTicker(userId int, input trade.BuySellTickerInput, price float64, count int) (float64, error)
	GetAllTickers(userId int) ([]trade.Portfolio, error)
	GetTickerByNasdaq(userId int, nasdaq string) (trade.Portfolio, error)
}

type Repository struct {
	Authorization
	Trade
	TypeTrade
	Portfolio
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		Trade:         NewTradePostgres(db),
		TypeTrade:     NewTypeTradePostgres(db),
		Portfolio:     NewPortfolioPostgres(db),
	}
}
