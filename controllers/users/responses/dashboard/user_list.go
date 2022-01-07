package dashboard

import "fgd-alterra-29/business/users"

type UserList struct {
	ID     int    `json:"user_id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Photo  string `json:"photo_url"`
	Status string `json:"status"`
}

func ToUserList(Domain users.Domain) UserList {
	return UserList{
		ID:     Domain.ID,
		Name:   Domain.Name,
		Email:  Domain.Email,
		Photo:  Domain.Photo_url,
		Status: Domain.Status,
	}
}

func ToListUserList(u []users.Domain) []UserList {
	var Domains []UserList

	for _, val := range u {
		Domains = append(Domains, ToUserList(val))
	}
	return Domains
}
