package trade

type User struct {
	Id             int    `json:"_"`
	Name           string `json:"name" binding:"required"`
	Username       string `json:"username" binding:"required"`
	Password       string `json:"password" binding:"required"`
	Balance        int    `json:"balance"`
	Trading_status string `json:"trading_status"`
}

type SignInUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}
