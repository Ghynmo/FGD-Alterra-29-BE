package request

import "fgd-alterra-29/business/users"

type UpdateProfile struct {
	Name    string `form:"name" json:"name"`
	Photo   string `form:"photo" json:"photo"`
	Email   string `form:"email" json:"email"`
	Phone   string `form:"phone" json:"phone"`
	Bio     string `form:"bio" json:"bio"`
	Address string `form:"address" json:"address"`
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
