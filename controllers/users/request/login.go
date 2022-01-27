package request

import "fgd-alterra-29/business/users"

type Login struct {
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (log *Login) FromLogin() users.Domain {
	return users.Domain{
		Email:    log.Email,
		Password: log.Password,
	}
}
