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
	Trade_type string `json:"trade_type"`
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
