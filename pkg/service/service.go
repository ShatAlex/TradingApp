package service

import (
	trade "github.com/ShatAlex/trading-app"
	"github.com/ShatAlex/trading-app/pkg/repository"
)

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
	GetAll(userId int) ([]trade.TypeTrade, error)
	GetTypeById(userId, typeId int) (trade.TypeTrade, error)
	Delete(userId, typeId int) error
	Update(userId, typeId int, typeTrade trade.TypeTrade) error
}

type Service struct {
	Authorization
	Trade
	TypeTrade
}

func NewService(rep *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(rep.Authorization),
		TypeTrade:     NewTypeTradeService(rep.TypeTrade),
		Trade:         NewTradeService(rep.Trade, rep.TypeTrade),
	}
}
