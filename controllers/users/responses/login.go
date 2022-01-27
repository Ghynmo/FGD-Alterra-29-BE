package responses

import "fgd-alterra-29/business/users"

type LoginResponse struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Token string `json:"token"`
}

func ToLoginResponse(domain users.Domain) LoginResponse {
	return LoginResponse{
		Name:  domain.Name,
		Email: domain.Email,
		Token: domain.Token,
	}
}
