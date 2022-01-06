package editprofile

import (
	"fgd-alterra-29/business/users"
)

type UserEdit struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Photo   string `json:"photo"`
	Phone   string `json:"phone"`
	Bio     string `json:"bio"`
	Address string `json:"address"`
}

func ToUserEdit(domain users.Domain) UserEdit {
	return UserEdit{
		Name:    domain.Name,
		Email:   domain.Email,
		Photo:   domain.Photo_url,
		Phone:   domain.Phone,
		Bio:     domain.Bio,
		Address: domain.Address,
	}
}
