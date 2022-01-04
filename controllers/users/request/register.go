package request

import "fgd-alterra-29/business/users"

type Register struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (reg *Register) FromRegister() users.Domain {
	return users.Domain{
		Name:     reg.Name,
		Email:    reg.Email,
		Password: reg.Password,
	}
}
