package request

import "fgd-alterra-29/business/users"

type Login struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (log *Login) FromLogin() users.Domain {
	return users.Domain{
		Email:    log.Email,
		Password: log.Password,
	}
}
