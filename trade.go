package trade

import "errors"

type Trade struct {
	Id      int    `json:"id" db:"id"`
	Ticker  string `json:"ticker" db:"ticker"`
	User_id int    `json:"user_id" db:"user_id"`
	Type_id int    `json:"type_id" db:"type_id"`
	Price   int    `json:"price" db:"price"`
	Amount  int    `json:"amount" db:"amount"`
}

type TypeTrade struct {
	Id         int    `json:"id"`
	UserId     int    `json:"user_id" db:"user_id"` // db т.к. создавали без userId в Body, а без него GET не сработает
	Trade_type string `json:"trade_type"`
}

type UpdateTradeInput struct {
	Ticker  *string `json:"ticker"`
	Type_id *int    `json:"type_id"`
	Price   *int    `json:"price"`
	Amount  *int    `json:"amount"`
}

func (i UpdateTradeInput) Validate() error {
	if i.Ticker == nil && i.Price == nil && i.Type_id == nil && i.Amount == nil {
		return errors.New("update structure has no values")
	}
	return nil
}

type PolygonInput struct {
	Ticker *string `json:"ticker"`
}
