package follows

import (
	"context"
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

func (uc FollowUseCase) GetFollowers(ctx context.Context, id int) ([]Domain, error) {
	followers, err := uc.Repo.GetFollowers(ctx, id)

	if err != nil {
		return []Domain{}, err
	}

	return followers, nil
}

func (uc FollowUseCase) GetFollowing(ctx context.Context, id int) ([]Domain, error) {
	following, err := uc.Repo.GetFollowing(ctx, id)

	if err != nil {
		return []Domain{}, err
	}

	return following, nil
}

func (uc FollowUseCase) FollowsController(ctx context.Context, domain Domain) (Domain, error) {
	following, err := uc.Repo.Follows(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return following, nil
}

func (uc FollowUseCase) UnfollowController(ctx context.Context, domain Domain) (Domain, error) {
	following, err := uc.Repo.Unfollow(ctx, domain)

	if err != nil {
		return Domain{}, err
	}

	return following, nil
}
