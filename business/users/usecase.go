package users

import (
	"context"
	"errors"
	"fgd-alterra-29/app/middlewares"
	"fgd-alterra-29/helpers"
	"time"
)

type UserUseCase struct {
	ConfigJwt      middlewares.ConfigJWT
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserUseCase(repo Repository, timeout time.Duration, ConfigJWT middlewares.ConfigJWT) UseCase {
	return &UserUseCase{
		ConfigJwt:      ConfigJWT,
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *UserUseCase) RegisterController(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Name == "" {
		return Domain{}, errors.New("NAME MUST BE FILLED")
	}
	if domain.Email == "" {
		return Domain{}, errors.New("EMAIL MUST BE FILLED")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("PASSWORD MUST BE FILLED")
	}

	var errHash error
	domain.Password, errHash = helpers.Hash(domain.Password)
	if errHash != nil {
		return Domain{}, errHash
	}

	user, err := uc.Repo.Register(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	var errT error
	user.Token, errT = uc.ConfigJwt.GenerateToken(user.ID)
	if errT != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) LoginController(ctx context.Context, domain Domain) (Domain, error) {
	if domain.Email == "" {
		return Domain{}, errors.New("EMAIL MUST BE FILLED")
	}
	if domain.Password == "" {
		return Domain{}, errors.New("PASSWORD MUST BE FILLED")
	}

	user, err := uc.Repo.Login(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	if !(helpers.ValidateHash(domain.Password, user.Password)) {
		return Domain{}, errors.New("WRONG PASSWORD")
	}

	var errT error
	user.Token, errT = uc.ConfigJwt.GenerateToken(user.ID)
	if errT != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUsersController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetUsers(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUsersByNameController(ctx context.Context, name string) ([]Domain, error) {
	user, err := uc.Repo.GetUsersByName(ctx, name)
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

func (uc *UserUseCase) UpdateProfile(ctx context.Context, domain Domain, id int) (Domain, error) {
	admin, err := uc.Repo.UpdateProfile(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return admin, nil
}

func (uc *UserUseCase) BannedUserController(ctx context.Context, id int) (Domain, error) {
	state, _ := uc.Repo.GetBannedState(ctx, id)

	if state.Status == "active" {
		threads, err := uc.Repo.BannedUser(ctx, id)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}

	if state.Status == "banned" {
		threads, err := uc.Repo.UnbannedUser(ctx, id)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}
	return Domain{}, nil
}
