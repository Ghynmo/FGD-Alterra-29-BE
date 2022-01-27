package threadlikes

import (
	"context"
	up "fgd-alterra-29/business/user_points"
	"time"
)

type ThreadLikeUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
	UserPointRepo  up.Repository
}

func NewThreadLikeUseCase(repo Repository, timeout time.Duration, up up.Repository) UseCase {
	return &ThreadLikeUseCase{
		Repo:           repo,
		contextTimeout: timeout,
		UserPointRepo:  up,
	}
}

func (uc *ThreadLikeUseCase) LikeController(ctx context.Context, domain Domain, id int) (Domain, error) {
	state, errState := uc.Repo.GetLikeState(ctx, domain, id)
	if errState != nil {
		return Domain{}, errState
	}

	if state.User_id == 0 {
		threads, t_user_id, err := uc.Repo.NewLike(ctx, domain, id)
		uc.UserPointRepo.AddReputationPoint(ctx, 2, t_user_id)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}

	if state.User_id != 0 && !state.State {
		threads, t_user_id, err := uc.Repo.Like(ctx, domain, id)
		uc.UserPointRepo.AddReputationPoint(ctx, 2, t_user_id)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}

	if state.User_id != 0 && state.State {
		threads, t_user_id, err := uc.Repo.Unlike(ctx, domain, id)
		uc.UserPointRepo.AddReputationPoint(ctx, -2, t_user_id)
		if err != nil {
			return Domain{}, err
		}
		return threads, nil
	}

	return Domain{}, nil
}
