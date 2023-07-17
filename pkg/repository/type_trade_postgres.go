package repository

import (
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

func (r *TypeTradePostgres) Create(input trade.TypeTrade) (int, error) {
	var typeId int

	createTypeTrade := fmt.Sprintf("INSERT INTO %s trade_type VALUES $2 RETURNING id", typesTable)
	row := r.db.QueryRow(createTypeTrade, input.Trade_type)

	err := row.Scan(&typeId)
	if err != nil {
		return 0, err
	}

	return typeId, nil
}

func (r *TypeTradePostgres) GetAll() ([]trade.TypeTrade, error) {
	var types []trade.TypeTrade

	quary := fmt.Sprintf("SELECT * from %s", typesTable)

	err := r.db.Select(&types, quary)

	return types, err
}

func (r *TypeTradePostgres) GetTypeById(typeId int) (trade.TypeTrade, error) {
	var item trade.TypeTrade

	quary := fmt.Sprintf("SELECT * FROM %s WHERE id = $1", typesTable)

	if err := r.db.Get(&item, quary, typeId); err != nil {
		return item, err
	}

	return item, nil
}

func (r *TypeTradePostgres) Delete(typeId int) error {
	quary := fmt.Sprintf("DELETE FROM %s WHERE id = $1", typesTable)

	_, err := r.db.Exec(quary, typeId)
	return err
}

func (r *TypeTradePostgres) Update(typeId int, input trade.TypeTrade) error {
	quary := fmt.Sprintf("UPDATE %s SET trade_type = $1 WHERE id = $2", typesTable)

	_, err := r.db.Exec(quary, input.Trade_type, typeId)
	return err

}

func (r *TypeTradePostgres) SuperUserValidate(userId int) (bool, error) {
	flag := false
	isSuperUser := fmt.Sprintf("SELECT true FROM %s WHERE id = $1 AND is_superuser = $2", usersTable)
	row := r.db.QueryRow(isSuperUser, userId, true)

	_ = row.Scan(&flag)

	return flag, nil
}
