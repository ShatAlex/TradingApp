package service

import (
	trade "github.com/ShatAlex/trading-app"
	"github.com/ShatAlex/trading-app/pkg/repository"
)

type PortfolioService struct {
	repPortfolio repository.Portfolio
}

func NewPortfolioService(repPortfolio repository.Portfolio) *PortfolioService {
	return &PortfolioService{repPortfolio: repPortfolio}
}

func (s *PortfolioService) BuyTicker(userId int, input trade.BuySellTickerInput, price float64) (int, error) {
	return s.repPortfolio.BuyTicker(userId, input, price)
}

func (s *PortfolioService) SellTicker(userId int, input trade.BuySellTickerInput, price float64) (float64, error) {

	ticker, err := s.repPortfolio.GetTickerByNasdaq(userId, *input.Ticker)
	if err != nil {
		// Нет бумаги в портфеле
		return 0, err
	}

	return s.repPortfolio.SellTicker(userId, input, price, ticker.Amount)
}

func (s *PortfolioService) GetAllTickers(userId int) ([]trade.Portfolio, error) {
	return s.repPortfolio.GetAllTickers(userId)
}

func (s *PortfolioService) GetTickerByNasdaq(userId int, nasdaq string) (trade.Portfolio, error) {
	return s.repPortfolio.GetTickerByNasdaq(userId, nasdaq)
}
