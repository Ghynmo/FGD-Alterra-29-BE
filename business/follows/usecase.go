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
