package service

import (
	trade "github.com/ShatAlex/trading-app"
	"github.com/ShatAlex/trading-app/pkg/repository"
)

//go:generate mockgen -source=service.go -destination=mocks/mock.go

type Authorization interface {
	CreateUser(user trade.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
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
	GetAll() ([]trade.TypeTrade, error)
	GetTypeById(typeId int) (trade.TypeTrade, error)
	Delete(userId, typeId int) error
	Update(userId, typeId int, typeTrade trade.TypeTrade) error
	SuperUserValidate(userId int) (bool, error)
}

type Portfolio interface {
	BuyTicker(userId int, input trade.BuySellTickerInput, price float64) (int, error)
	SellTicker(userId int, input trade.BuySellTickerInput, price float64) (float64, error)
	GetAllTickers(userId int) ([]trade.Portfolio, error)
	GetTickerByNasdaq(userId int, nasdaq string) (trade.Portfolio, error)
}

type Service struct {
	Authorization
	Trade
	TypeTrade
	Portfolio
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
		Trade:         NewTradeService(rep.Trade, rep.TypeTrade, rep.Portfolio),
		TypeTrade:     NewTypeTradeService(rep.TypeTrade, rep.Portfolio),
		Portfolio:     NewPortfolioService(rep.Portfolio),
	}
}
