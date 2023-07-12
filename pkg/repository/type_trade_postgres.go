package repository

import (
	"errors"
	"fmt"

	trade "github.com/ShatAlex/trading-app"
	"github.com/jmoiron/sqlx"
)

type TypeTradePostgres struct {
	db *sqlx.DB
}

func NewTypeTradePostgres(db *sqlx.DB) *TypeTradePostgres {
	return &TypeTradePostgres{db: db}
}

func (r *TypeTradePostgres) Create(userId int, typeTrade trade.TypeTrade) (int, error) {
	var typeId int

	createTypeTrade := fmt.Sprintf("INSERT INTO %s (user_id, trade_type) VALUES ($1, $2) RETURNING id", typesTable)
	row := r.db.QueryRow(createTypeTrade, userId, typeTrade.Trade_type)

	err := row.Scan(&typeId)
	if err != nil {
		return 0, err
	}

	return typeId, nil
}

func (r *TypeTradePostgres) GetAll(userId int) ([]trade.TypeTrade, error) {
	var types []trade.TypeTrade

	quary := fmt.Sprintf("SELECT * from %s WHERE user_id = $1", typesTable)

	err := r.db.Select(&types, quary, userId)

	return types, err
}

func (r *TypeTradePostgres) GetTypeById(userId, typeId int) (trade.TypeTrade, error) {
	var item trade.TypeTrade

	quary := fmt.Sprintf("SELECT * FROM %s WHERE user_id = $1 and id = $2", typesTable)

	if err := r.db.Get(&item, quary, userId, typeId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TypeTradePostgres) Delete(userId, typeId int) error {
	quary := fmt.Sprintf("DELETE FROM %s WHERE user_id = $1 and id = $2", typesTable)
	_, err := r.db.Exec(quary, userId, typeId)
	_, newErr := r.GetTypeById(userId, typeId)
	if newErr != nil {
		return errors.New("persmissions denied")
	}
	return err
}

func (r *TypeTradePostgres) Update(userId, typeId int, typeTrade trade.TypeTrade) error {
	quary := fmt.Sprintf("UPDATE %s SET trade_type = $1 WHERE user_id = $2 and id = $3", typesTable)
	_, err := r.db.Exec(quary, typeTrade.Trade_type, userId, typeId)
	_, newErr := r.GetTypeById(userId, typeId)
	if newErr != nil {
		return errors.New("persmissions denied")
	}
	return err

}
