package badges

import (
	"context"
	"time"
)

type BadgeUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewBadgeUseCase(repo Repository, timeout time.Duration) UseCase {
	return &BadgeUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *BadgeUseCase) GetBadgesByUserController(ctx context.Context, id int) ([]Domain, error) {
	badge, err := uc.Repo.GetBadgesByUser(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return badge, nil
}
