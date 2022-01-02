package settings

import (
	"fgd-alterra-29/business/users"
)

type AdminSetting struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Photo string `json:"photo"`
	Phone string `json:"phone"`
}

func ToAdminSetting(domain users.Domain) AdminSetting {
	return AdminSetting{
		Name:  domain.Name,
		Email: domain.Email,
		Photo: domain.Photo_url,
		Phone: domain.Phone,
	}
}
