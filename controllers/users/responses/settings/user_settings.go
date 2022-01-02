package settings

import (
	"fgd-alterra-29/business/users"
)

type UserSetting struct {
	Name    string `json:"name"`
	Email   string `json:"email"`
	Photo   string `json:"photo"`
	Phone   string `json:"phone"`
	Bio     string `json:"bio"`
	Address string `json:"address"`
}

func ToUserSetting(domain users.Domain) UserSetting {
	return UserSetting{
		Name:    domain.Name,
		Email:   domain.Email,
		Photo:   domain.Photo_url,
		Phone:   domain.Phone,
		Bio:     domain.Bio,
		Address: domain.Address,
	}
}
