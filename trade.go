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
	UserId     int    `json:"user_id" db:"user_id"` // db т.к. создавали без userId в Body, а без него GET не сработает
	Trade_type string `json:"trade_type"`
}
