package repository

import (
	"fmt"
	"strings"

	trade "github.com/ShatAlex/trading-app"
	"github.com/jmoiron/sqlx"
)

type TradePostgres struct {
	db *sqlx.DB
}

func NewTradePostgres(db *sqlx.DB) *TradePostgres {
	return &TradePostgres{db: db}
}

func (r *TradePostgres) Create(userId int, trade trade.Trade) (int, error) {

	var tradeId int
	// var buyId, sellId int

	// select_type := fmt.Sprintf("SELECT id FROM %s WHERE trade_type = 'Покупка ценных бумаг'", typesTable)
	// if err := r.db.Get(&buyId, select_type); err != nil {
	// 	return 0, err
	// }

	// select_type = fmt.Sprintf("SELECT id FROM %s WHERE trade_type = 'Продажа ценных бумаг'", typesTable)
	// if err := r.db.Get(&sellId, select_type); err != nil {
	// 	return 0, err
	// }

	// if buyId == trade.TypeId {
	// 	return 0, errors.New("bad request. Try /api/v1/portfolio/buy")
	// }

	// if sellId == trade.TypeId {
	// 	return 0, errors.New("bad request. Try /api/v1/portfolio/sell")
	// }

	createTrade := fmt.Sprintf(`INSERT INTO %s (ticker, user_id, type_id, price, amount) SELECT $1, $2, $3, $4, $5
								WHERE (SELECT true FROM %s ty WHERE ty.id = $6)
								RETURNING id`, tradesTable, typesTable)
	row := r.db.QueryRow(createTrade, trade.Ticker, userId, trade.TypeId, trade.Price, trade.Amount, trade.TypeId)

	err := row.Scan(&tradeId)
	if err != nil {
		return 0, err
	}

	return tradeId, nil
}

func (r *TradePostgres) GetAll(userId int) ([]trade.Trade, error) {
	var trades []trade.Trade

	quary := fmt.Sprintf("SELECT * from %s WHERE user_id = $1", tradesTable)

	err := r.db.Select(&trades, quary, userId)

	return trades, err
}

func (r *TradePostgres) GetTradeById(userId, tradeId int) (trade.Trade, error) {
	var item trade.Trade

	quary := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 and id = $2", tradesTable)

	if err := r.db.Get(&item, quary, userId, tradeId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TradePostgres) Delete(userId, tradeId int) error {
	quary := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 and id = $2", tradesTable)
	_, err := r.db.Exec(quary, userId, tradeId)
	return err
}

func (r *TradePostgres) Update(userId, tradeId int, trade trade.UpdateTradeInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1
	query := ""

	if trade.Ticker != nil {
		setValues = append(setValues, fmt.Sprintf("ticker=$%d", argId))
		args = append(args, trade.Ticker)
		argId++
	}

	if trade.Price != nil {
		setValues = append(setValues, fmt.Sprintf("price=$%d", argId))
		args = append(args, trade.Price)
		argId++
	}

	if trade.Amount != nil {
		setValues = append(setValues, fmt.Sprintf("amount=$%d", argId))
		args = append(args, trade.Amount)
		argId++
	}
	if trade.Typeid != nil {
		setValues = append(setValues, fmt.Sprintf("type_id=$%d", argId))
		args = append(args, trade.Typeid)
		argId++

		setQuery := strings.Join(setValues, ", ")
		query = fmt.Sprintf(`UPDATE %s tr SET %s FROM %s ty WHERE tr.id = $%d AND tr.user_id = $%d
			AND ty.id = $%d AND ty.user_id = $%d`, tradesTable, setQuery, typesTable, argId, argId+1, argId+2, argId+3)
		args = append(args, tradeId, userId, trade.Typeid, userId)

	} else {
		setQuery := strings.Join(setValues, ", ")
		query = fmt.Sprintf("UPDATE %s SET %s WHERE id = $%d AND user_id = $%d",
			tradesTable, setQuery, argId, argId+1)
		args = append(args, tradeId, userId)

	}
	_, err := r.db.Exec(query, args...)

	return err
}
