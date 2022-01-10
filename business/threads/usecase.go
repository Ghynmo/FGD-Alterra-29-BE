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

func (uc *ThreadUseCase) GetThreadsByTitleController(ctx context.Context, title string) ([]Domain, error) {
	thread, err := uc.Repo.GetThreadsByTitle(ctx, title)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetProfileThreads(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetProfileThreads(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetThreadQuantity(ctx context.Context) (Domain, error) {
	thread, err := uc.Repo.GetThreadQuantity(ctx)
	if err != nil {
		return Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetHomepageThreads(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetHomepageThreads(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetThreads(ctx context.Context) ([]Domain, error) {
	thread, err := uc.Repo.GetThreads(ctx)
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

func (uc *ThreadUseCase) DeleteThread(ctx context.Context, id int) (Domain, error) {
	thread, err := uc.Repo.DeleteThread(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetHotThreads(ctx context.Context) ([]Domain, error) {
	thread, err := uc.Repo.GetHotThreads(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetSearch(ctx context.Context, threadname string) ([]Domain, error) {
	thread, err := uc.Repo.GetSearch(ctx, threadname)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}
