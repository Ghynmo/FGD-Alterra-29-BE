package request

import "fgd-alterra-29/business/users"

type Register struct {
	Name     string `form:"name" json:"name"`
	Email    string `form:"email" json:"email"`
	Password string `form:"password" json:"password"`
}

func (reg *Register) FromRegister() users.Domain {
	return users.Domain{
		Name:     reg.Name,
		Email:    reg.Email,
		Password: reg.Password,
	}
}
