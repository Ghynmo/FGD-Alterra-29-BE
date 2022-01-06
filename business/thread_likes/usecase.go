package threadlikes

import (
	"context"
	"time"
)

type ThreadLikeUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewThreadLikeUseCase(repo Repository, timeout time.Duration) UseCase {
	return &ThreadLikeUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *ThreadLikeUseCase) LikeController(ctx context.Context, domain Domain) (Domain, error) {
	threadlikes, err := uc.Repo.Like(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return threadlikes, nil
}

func (uc *ThreadLikeUseCase) UnlikeController(ctx context.Context, domain Domain) (Domain, error) {
	threadlikes, err := uc.Repo.Unlike(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return threadlikes, nil
}
