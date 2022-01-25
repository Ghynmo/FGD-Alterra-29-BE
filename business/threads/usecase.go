package threads

import (
	"context"
	userpoints "fgd-alterra-29/business/user_points"
	"time"
)

type ThreadUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
	RepoPoint      userpoints.Repository
}

func NewThreadUseCase(repo Repository, timeout time.Duration, up userpoints.Repository) UseCase {
	return &ThreadUseCase{
		Repo:           repo,
		contextTimeout: timeout,
		RepoPoint:      up,
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

func (uc *ThreadUseCase) CreateThread(ctx context.Context, domain Domain, id int) (Domain, error) {
	thread, err := uc.Repo.CreateThread(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	_, err2 := uc.RepoPoint.AddThreadPoint(ctx, domain.User_id)
	if err2 != nil {
		return Domain{}, err2
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

func (uc *ThreadUseCase) GetThreadByID(ctx context.Context, id int) (Domain, error) {
	thread, err := uc.Repo.GetThreadByID(ctx, id)
	if err != nil {
		return Domain{}, err
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

func (uc *ThreadUseCase) ActivateThread(ctx context.Context, id int) (Domain, error) {
	thread, err := uc.Repo.ActivateThread(ctx, id)
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

func (uc *ThreadUseCase) GetThreadsByCategoryID(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetThreadsByCategoryID(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetSideNewsThreads(ctx context.Context) ([]Domain, error) {
	thread, err := uc.Repo.GetSideNewsThreads(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}
