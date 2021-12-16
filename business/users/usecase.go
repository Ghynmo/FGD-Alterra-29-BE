package users

import (
	"context"
	"time"
)

type UserUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUseCase(repo Repository, timeout time.Duration) UseCase {
	return &UserUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *UserUseCase) GetUserController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetUser(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}
