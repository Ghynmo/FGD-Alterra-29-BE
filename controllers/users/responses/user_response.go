package responses

import (
	"fgd-alterra-29/business/users"
)

type UserResponse struct {
	Name string `json:"name"`
}

func FromDomain(domain users.Domain) UserResponse {
	return UserResponse{
		Name: domain.Name,
	}
}

func FromListDomain(data []users.Domain) (result []UserResponse) {
	result = []UserResponse{}
	for _, val := range data {
		result = append(result, FromDomain(val))
	}
	return result
}
