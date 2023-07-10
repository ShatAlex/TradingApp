package trade

type Trade struct {
	Id      int    `json:"id"`
	Figi    string `json:"figi"`
	User_id int    `json:"user_id"`
	Type_id int    `json:"type_id"`
	Price   int    `json:"price"`
	Amount  int    `json:"amount"`
}

type TypeTrade struct {
	Id         int    `json:"id"`
	Trade_type string `json:"trade_type"`
}
