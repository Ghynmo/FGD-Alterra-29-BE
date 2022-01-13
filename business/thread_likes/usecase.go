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
	state, _ := uc.Repo.GetLikeState(ctx, domain)

	if state.User_id == 0 {
		threads, err := uc.Repo.NewLike(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}

	if state.User_id != 0 && !state.State {
		threads, err := uc.Repo.Like(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}

	if state.User_id != 0 && state.State {
		threads, err := uc.Repo.Unlike(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}

	return Domain{}, nil
}
