package repository

import (
	"errors"
	"fmt"

	trade "github.com/ShatAlex/trading-app"
	"github.com/jmoiron/sqlx"
)

type PortfolioPostgres struct {
	db *sqlx.DB
}

func NewPortfolioPostgres(db *sqlx.DB) *PortfolioPostgres {
	return &PortfolioPostgres{db: db}
}

func (r *PortfolioPostgres) BuyTicker(userId int, input trade.BuySellTickerInput, price float64) (int, error) {

	var tradeId, typeId, portfolioID int

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	select_type := fmt.Sprintf("SELECT id FROM %s WHERE trade_type = 'Покупка ценных бумаг'", typesTable)
	if err := r.db.Get(&typeId, select_type); err != nil {
		tx.Rollback()
		return 0, err
	}

	createTrade := fmt.Sprintf(`INSERT INTO %s (ticker, user_id, type_id, price, amount) 
								VALUES ($1, $2, $3, $4, $5) RETURNING id`, tradesTable)
	row := r.db.QueryRow(createTrade, input.Ticker, userId, typeId, price, input.Amount)

	err = row.Scan(&tradeId)
	if err != nil {
		tx.Rollback()
		return 0, err
	}

	addToPortfolio := fmt.Sprintf(`INSERT INTO %s (user_id, ticker, amount) SELECT $1, $2, $3
								   WHERE NOT EXISTS (SELECT true FROM %s WHERE user_id = $4 AND ticker = $5) RETURNING id`,
		portfolioTable, portfolioTable)
	row = r.db.QueryRow(addToPortfolio, userId, input.Ticker, input.Amount, userId, input.Ticker)
	err = row.Scan(&portfolioID)
	if err != nil {

		var amount *int

		exist := fmt.Sprintf("SELECT amount FROM %s WHERE user_id = $1 AND ticker = $2", portfolioTable)
		if err = r.db.Get(&amount, exist, userId, input.Ticker); err != nil {
			return 0, err
		}

		updateToPortfolio := fmt.Sprintf("UPDATE %s SET amount = $1 WHERE user_id = $2 AND ticker = $3", portfolioTable)
		_, err = r.db.Exec(updateToPortfolio, *input.Amount+*amount, userId, input.Ticker)
		if err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	return tradeId, tx.Commit()
}

func (r *PortfolioPostgres) SellTicker(userId int, input trade.BuySellTickerInput, price float64, count int) (float64, error) {
	var total float64

	tx, err := r.db.Begin()
	if err != nil {
		return 0, err
	}

	totalCount := count - *input.Amount

	if totalCount < 0 {
		return 0, errors.New("there are not so many tickers")
	} else if totalCount == 0 {
		deleteTicker := fmt.Sprintf("DELETE FROM %s WHERE ticker = $1 AND user_id = $2", portfolioTable)
		if _, err := r.db.Exec(deleteTicker, input.Ticker, userId); err != nil {
			tx.Rollback()
			return 0, err
		}
	} else {
		updateTicker := fmt.Sprintf("UPDATE %s SET ticker = $1, user_id = $2, amount = $3 WHERE ticker = $4 AND user_id = $5",
			portfolioTable)
		if _, err := r.db.Exec(updateTicker, input.Ticker, userId, totalCount, input.Ticker, userId); err != nil {
			tx.Rollback()
			return 0, err
		}
	}

	total = float64(*input.Amount) * price

	return total, tx.Commit()

}

func (r *PortfolioPostgres) GetAllTickers(userId int) ([]trade.Portfolio, error) {
	var tickers []trade.Portfolio

	quary := fmt.Sprintf("SELECT * from %s WHERE user_id = $1", portfolioTable)

	err := r.db.Select(&tickers, quary, userId)

	return tickers, err
}

func (r *PortfolioPostgres) GetTickerByNasdaq(userId int, nasdaq string) (trade.Portfolio, error) {
	var ticker trade.Portfolio

	quary := fmt.Sprintf("SELECT * from %s WHERE user_id = $1 AND ticker = $2", portfolioTable)

	err := r.db.Get(&ticker, quary, userId, nasdaq)

	return ticker, err
}
