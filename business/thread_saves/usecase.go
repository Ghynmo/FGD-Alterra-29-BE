package threadsaves

import (
	"context"
	"time"
)

type ThreadSaveUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewThreadSaveUseCase(repo Repository, timeout time.Duration) UseCase {
	return &ThreadSaveUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *ThreadSaveUseCase) SaveThreadController(ctx context.Context, domain Domain) (Domain, error) {
	threadsave, err := uc.Repo.SaveThread(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return threadsave, nil
}

func (uc *ThreadSaveUseCase) UnsaveThreadController(ctx context.Context, domain Domain) (Domain, error) {
	threadsave, err := uc.Repo.UnsaveThread(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return threadsave, nil
}
