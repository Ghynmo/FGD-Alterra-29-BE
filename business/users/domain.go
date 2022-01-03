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
	LoginController(ctx context.Context, domain Domain) (Domain, error)
	GetUserController(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Login(ctx context.Context, domain Domain) (Domain, error)
	GetUser(ctx context.Context) ([]Domain, error)
}
