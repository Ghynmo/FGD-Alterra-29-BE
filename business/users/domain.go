package users

import (
	"context"
)

type Domain struct {
	ID       int
	Name     string
	Password string
}

type UseCase interface {
	GetUserController(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	GetUser(ctx context.Context) ([]Domain, error)
}
