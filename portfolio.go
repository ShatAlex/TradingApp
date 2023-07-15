package trade

type Portfolio struct {
	Id     int    `json:"id"`
	Userid int    `json:"user_id" db:"user_id"`
	Ticker string `json:"ticker"`
	Amount int    `json:"amount"`
}

type BuySellTickerInput struct {
	Ticker *string `json:"ticker"`
	Amount *int    `json:"amount"`
}
