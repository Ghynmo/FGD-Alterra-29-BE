package threads

import (
	"context"
	"time"
)

type ThreadUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewThreadUseCase(repo Repository, timeout time.Duration) UseCase {
	return &ThreadUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *ThreadUseCase) GetProfileThreads(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetProfileThreads(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetRecommendationThreads(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetRecommendationThreads(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetHotThreads(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetHotThreads(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}
