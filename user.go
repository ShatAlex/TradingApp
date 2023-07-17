package trade

import "errors"

type User struct {
	Id          int     `json:"_"`
	Name        string  `json:"name" binding:"required"`
	Username    string  `json:"username" binding:"required"`
	Password    string  `json:"password" binding:"required"`
	Balance     int     `json:"balance"`
	IsSuperUser *string `json:"is_superuser"`
}

type SignInUser struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (u User) SuperUserValidate() error {
	if u.IsSuperUser != nil {
		return errors.New("bad request. Field is_superuser is not editable")
	}
	return nil
}
