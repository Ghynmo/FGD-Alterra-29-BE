package request

import "fgd-alterra-29/business/users"

type UpdateProfile struct {
	Name    string `form:"name"`
	Photo   string `form:"photo"`
	Email   string `form:"email"`
	Phone   string `form:"phone"`
	Bio     string `form:"bio"`
	Address string `form:"address"`
}

func (us *UpdateProfile) ToDomain() users.Domain {
	return users.Domain{
		Name:      us.Name,
		Photo_url: us.Photo,
		Email:     us.Email,
		Phone:     us.Phone,
		Bio:       us.Bio,
		Address:   us.Address,
	}
}
