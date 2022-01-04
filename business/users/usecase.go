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
	user.Token, errT = uc.ConfigJwt.GenerateToken(domain.ID)
	if errT != nil {
		return Domain{}, err
	}

	return user, nil
}

func (uc *UserUseCase) GetUserController(ctx context.Context) ([]Domain, error) {
	user, err := uc.Repo.GetUser(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return user, nil
}
