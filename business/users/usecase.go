package users

import (
	"context"
	"fgd-alterra-29/business/threads"
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

func (uc *UserUseCase) GetUsersController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetUsers(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetProfileController(ctx context.Context, id int) (Domain, []threads.Domain, error) {
	user, err := uc.Repo.GetProfile(ctx, id)
	if err != nil {
		return Domain{}, []threads.Domain{}, err
	}

	threads, _ := uc.Repo.GetCategoryActive(ctx, id)

	return user, threads, nil
}
