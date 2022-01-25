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

func (uc *ThreadSaveUseCase) SaveThreadController(ctx context.Context, domain Domain, id int) (Domain, error) {
	state, errGet := uc.Repo.GetSaveState(ctx, domain, id)
	if errGet != nil {
		return Domain{}, errGet
	}

	if state.User_id == 0 {
		threads, err := uc.Repo.NewSave(ctx, domain, id)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}

	if state.User_id != 0 && !state.State {
		threads, err := uc.Repo.Save(ctx, domain, id)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}

	if state.User_id != 0 && state.State {
		threads, err := uc.Repo.Unsave(ctx, domain, id)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}

	return Domain{}, nil
}
