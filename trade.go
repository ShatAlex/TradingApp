package trade

import "errors"

type Trade struct {
	Id     int     `json:"id" db:"id"`
	Ticker string  `json:"ticker" db:"ticker"`
	UserId int     `json:"user_id" db:"user_id"`
	TypeId int     `json:"type_id" db:"type_id"`
	Price  float64 `json:"price" db:"price"`
	Amount int     `json:"amount" db:"amount"`
}

type TypeTrade struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id" db:"user_id"` // db т.к. создавали без userId в Body, а без него GET не сработает
	Trade_type string `json:"trade_type"`
}

type Portfolio struct {
	Id     int    `json:"id"`
	Userid int    `json:"user_id" db:"user_id"`
	Ticker string `json:"ticker"`
	Count  int    `json:"count"`
}

type UpdateTradeInput struct {
	Ticker *string  `json:"ticker"`
	Typeid *int     `json:"type_id" db:"type_id"`
	Price  *float64 `json:"price"`
	Amount *int     `json:"amount"`
}

func (i UpdateTradeInput) Validate() error {
	if i.Ticker == nil && i.Price == nil && i.Typeid == nil && i.Amount == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type PolygonInput struct {
	Ticker *string `json:"ticker"`
}

type BuySellTickerInput struct {
	Ticker *string `json:"ticker"`
	Amount *int    `json:"amount"`
}
