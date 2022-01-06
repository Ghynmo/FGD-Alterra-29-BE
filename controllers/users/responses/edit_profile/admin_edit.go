package editprofile

import (
	"fgd-alterra-29/business/users"
)

type AdminEdit struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Photo string `json:"photo"`
	Phone string `json:"phone"`
}

func ToAdminEdit(domain users.Domain) AdminEdit {
	return AdminEdit{
		Name:  domain.Name,
		Email: domain.Email,
		Photo: domain.Photo_url,
		Phone: domain.Phone,
	}
}
