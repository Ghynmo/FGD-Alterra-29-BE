package threadshares

import (
	"context"
	"time"
)

type ThreadShareUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewThreadShareUseCase(repo Repository, timeout time.Duration) UseCase {
	return &ThreadShareUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *ThreadShareUseCase) ThreadShareController(ctx context.Context, domain Domain, id int) (Domain, error) {
	thread, err := uc.Repo.ThreadShare(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	return thread, nil
}
