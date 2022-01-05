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

func (uc *ThreadUseCase) GetProfileThreads(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetProfileThreads(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) CreateThread(ctx context.Context, domain Domain) (Domain, error) {
	thread, err := uc.Repo.CreateThread(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	uc.RepoPoint.AddThreadPoint(ctx, domain.User_id)

	return thread, nil
}
