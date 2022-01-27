package userpoints

import (
	"context"
	"time"
)

type UserPointUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserPointUseCase(repo Repository, timeout time.Duration) UseCase {
	return &UserPointUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *UserPointUseCase) AddThreadPointController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.AddThreadPoint(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserPointUseCase) AddPostPointController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.AddPostPoint(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserPointUseCase) AddReputationPointController(ctx context.Context, multiple int, id int) (Domain, error) {
	user, err := uc.Repo.AddReputationPoint(ctx, multiple, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}
