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

func (uc *UserUseCase) GetUsersController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetUsers(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetProfileController(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetProfile(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUsersQuantity(ctx context.Context) (Domain, error) {
	user, err := uc.Repo.GetUsersQuantity(ctx)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetProfileSetting(ctx context.Context, id int) (Domain, error) {
	user, err := uc.Repo.GetProfileSetting(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) UpdateSetting(ctx context.Context, domain Domain, id int) (Domain, error) {
	admin, err := uc.Repo.UpdateSetting(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return admin, nil
}
