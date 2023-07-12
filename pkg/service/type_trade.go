package service

import (
	trade "github.com/ShatAlex/trading-app"
	"github.com/ShatAlex/trading-app/pkg/repository"
)

type TypeTradeService struct {
	repType repository.TypeTrade
}

func NewTypeTradeService(repType repository.TypeTrade) *TypeTradeService {
	return &TypeTradeService{repType: repType}
}

func (s *TypeTradeService) Create(userId int, typeTrade trade.TypeTrade) (int, error) {
	return s.repType.Create(userId, typeTrade)
}

func (s *TypeTradeService) GetAll(userId int) ([]trade.TypeTrade, error) {
	return s.repType.GetAll(userId)
}

func (s *TypeTradeService) GetTypeById(userId, typeId int) (trade.TypeTrade, error) {
	return s.repType.GetTypeById(userId, typeId)
}

func (s *TypeTradeService) Delete(userId, typeId int) error {
	return s.repType.Delete(userId, typeId)
}

func (s *TypeTradeService) Update(userId, typeId int, typeTrade trade.TypeTrade) error {
	return s.repType.Update(userId, typeId, typeTrade)
}
