package responses

import (
	"fgd-alterra-29/business/users"
)

type UserResponse struct {
	Name   string `json:"name"`
	Email  string `json:"email"`
	Photo  string `json:"photo"`
	Status string `json:"status"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Name:   domain.Name,
		Email:  domain.Email,
		Photo:  domain.Photo_url,
		Status: domain.Status,
	}
}

func FromListDomain(data []users.Domain) (result []UserResponse) {
	result = []UserResponse{}
	for _, val := range data {
		result = append(result, FromDomain(val))
	}
	return result
}
