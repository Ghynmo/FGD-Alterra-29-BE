package userbadges

import (
	"context"
	"time"
)

type UserBadgeUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewUserBadgeUseCase(repo Repository, timeout time.Duration) UseCase {
	return &UserBadgeUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *UserBadgeUseCase) GetUserBadge(ctx context.Context, id int) ([]Domain, error) {
	userbadge, err := uc.Repo.GetUserBadge(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return userbadge, nil
}
