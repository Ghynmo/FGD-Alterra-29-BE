package responses

import "fgd-alterra-29/business/users"

type LoginResponse struct {
	Name  string
	Email string
	Token string
}

func ToLoginResponse(domain users.Domain) LoginResponse {
	return LoginResponse{
		Name:  domain.Name,
		Email: domain.Email,
		Token: domain.Token,
	}
}
