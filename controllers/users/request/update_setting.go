package request

import "fgd-alterra-29/business/users"

type UpdateSetting struct {
	Name  string `form:"name"`
	Photo string `form:"photo"`
	Email string `form:"email"`
	Phone string `form:"phone"`
}

func (us *UpdateSetting) ToDomain() users.Domain {
	return users.Domain{
		Name:      us.Name,
		Photo_url: us.Photo,
		Email:     us.Email,
		Phone:     us.Phone,
	}
}
