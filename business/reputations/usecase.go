package reputations

import (
	"context"
	"time"
)

type ReputationUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewReputationUseCase(repo Repository, timeout time.Duration) UseCase {
	return &ReputationUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *ReputationUseCase) CreateReputationController(ctx context.Context, domain Domain) (Domain, error) {

	reputation, err := uc.Repo.CreateReputation(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return reputation, nil
}

func (uc *ReputationUseCase) GetReputationByUser(ctx context.Context, id int) (Domain, error) {

	reputation, err := uc.Repo.GetReputationByUser(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return reputation, nil
}
