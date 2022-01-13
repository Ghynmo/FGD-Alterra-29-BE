package follows

import (
	"context"
	"errors"
	"time"
)

type FollowUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewFollowUseCase(repo Repository, time time.Duration) UseCase {
	return &FollowUseCase{
		Repo:           repo,
		contextTimeout: time,
	}
}

func (uc *FollowUseCase) GetFollowers(ctx context.Context, id int) ([]Domain, error) {
	followers, err := uc.Repo.GetFollowers(ctx, id)

	if err != nil {
		return []Domain{}, err
	}

	return followers, nil
}

func (uc *FollowUseCase) GetFollowing(ctx context.Context, id int) ([]Domain, error) {
	following, err := uc.Repo.GetFollowing(ctx, id)

	if err != nil {
		return []Domain{}, err
	}

	return following, nil
}

func (uc *FollowUseCase) FollowsController(ctx context.Context, domain Domain) (Domain, error) {

	if domain.User_id == domain.Follower_id {
		return Domain{}, errors.New("CANNOT FOLLOWS YOURSELF")
	}

	state, _ := uc.Repo.GetFollowState(ctx, domain)

	if state.User_id == 0 {
		follow, err := uc.Repo.NewFollow(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		return follow, nil
	}

	if state.User_id != 0 && !state.State {
		follow, err := uc.Repo.Follows(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		return follow, nil
	}

	if state.User_id != 0 && state.State {
		follow, err := uc.Repo.Unfollow(ctx, domain)
		if err != nil {
			return Domain{}, err
		}
		return follow, nil
	}

	return Domain{}, nil
}
