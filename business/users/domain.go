package users

import (
	"context"
)

type Domain struct {
	ID       int
	Name     string
	Email    string
	Password string
	Token    string
}

type UseCase interface {
	RegisterController(ctx context.Context, domain Domain) (Domain, error)
	LoginController(ctx context.Context, domain Domain) (Domain, error)
	GetUserController(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Register(ctx context.Context, domain Domain) (Domain, error)
	Login(ctx context.Context, domain Domain) (Domain, error)
	GetUser(ctx context.Context) ([]Domain, error)
}
