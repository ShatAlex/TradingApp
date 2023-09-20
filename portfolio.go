package trade

type Portfolio struct {
	Id     int    `json:"id"`
	Userid int    `json:"user_id" db:"user_id" binding:"required"`
	Ticker string `json:"ticker"`
	Amount int    `json:"amount"`
}

type BuySellTickerInput struct {
	Ticker *string `json:"ticker" binding:"required"`
	Amount *int    `json:"amount" binding:"required"`
}
