package service

import (
	trade "github.com/ShatAlex/trading-app"
	"github.com/ShatAlex/trading-app/pkg/repository"
)

type TradeService struct {
	repTrade     repository.Trade
	repTypeTrade repository.TypeTrade
}

func NewTradeService(repTrade repository.Trade, repTypeTrade repository.TypeTrade) *TradeService {
	return &TradeService{repTrade: repTrade, repTypeTrade: repTypeTrade}
}

func (s *TradeService) Create(userId int, trade trade.Trade) (int, error) {
	return s.repTrade.Create(userId, trade)
}

func (s *TradeService) GetAll(userId int) ([]trade.Trade, error) {
	return s.repTrade.GetAll(userId)
}

func (s *TradeService) GetTradeById(userId, tradeId int) (trade.Trade, error) {
	return s.repTrade.GetTradeById(userId, tradeId)
}

func (s *TradeService) Delete(userId, tradeId int) error {
	return s.repTrade.Delete(userId, tradeId)
}

func (s *TradeService) Update(userId, tradeId int, trade trade.UpdateTradeInput) error {
	if err := trade.Validate(); err != nil {
		return err
	}
	return s.repTrade.Update(userId, tradeId, trade)
}

func (s *TradeService) BuyTicker(userId int, input trade.BuySellTickerInput, price float64) (int, error) {
	return s.repTrade.BuyTicker(userId, input, price)
}

func (s *TradeService) SellTicker(userId int, input trade.BuySellTickerInput, price float64) (int, error) {
	return s.repTrade.SellTicker(userId, input, price)
}

func (s *TradeService) GetAllTickers(userId int) ([]trade.Portfolio, error) {
	return s.repTrade.GetAllTickers(userId)
}
