package apinews

import (
	"context"
	"time"
)

type APINewsUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewAPINewsUseCase(repo Repository, timeout time.Duration) UseCase {
	return &APINewsUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *APINewsUseCase) GetAPINews(ctx context.Context, apikey string) (Domain, error) {
	badge, err := uc.Repo.GetAPINews(ctx, apikey)
	if err != nil {
		return Domain{}, err
	}

	return badge, nil
}
