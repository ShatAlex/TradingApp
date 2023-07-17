package service

import (
	"errors"

	trade "github.com/ShatAlex/trading-app"
	"github.com/ShatAlex/trading-app/pkg/repository"
)

type TypeTradeService struct {
	repType      repository.TypeTrade
	repPortfolio repository.Portfolio
}

func NewTypeTradeService(repType repository.TypeTrade, repPortfolio repository.Portfolio) *TypeTradeService {
	return &TypeTradeService{repType: repType, repPortfolio: repPortfolio}
}

func (s *TypeTradeService) Create(userId int, typeTrade trade.TypeTrade) (int, error) {
	isSuperUser, err := s.repType.SuperUserValidate(userId)
	if err != nil {
		return 0, err
	}
	if !isSuperUser {
		return 0, errors.New("this method is available only to the administrator")
	}
	return s.repType.Create(typeTrade)
}

func (s *TypeTradeService) GetAll() ([]trade.TypeTrade, error) {
	return s.repType.GetAll()
}

func (s *TypeTradeService) GetTypeById(typeId int) (trade.TypeTrade, error) {
	return s.repType.GetTypeById(typeId)
}

func (s *TypeTradeService) Delete(userId, typeId int) error {
	isSuperUser, err := s.repType.SuperUserValidate(userId)
	if err != nil {
		return err
	}
	if !isSuperUser {
		return errors.New("this method is available only to the administrator")
	}
	return s.repType.Delete(typeId)
}

func (s *TypeTradeService) Update(userId, typeId int, typeTrade trade.TypeTrade) error {
	isSuperUser, err := s.repType.SuperUserValidate(userId)
	if err != nil {
		return err
	}
	if !isSuperUser {
		return errors.New("this method is available only to the administrator")
	}
	return s.repType.Update(typeId, typeTrade)
}

func (s *TypeTradeService) SuperUserValidate(userId int) (bool, error) {
	return s.repType.SuperUserValidate(userId)
}
