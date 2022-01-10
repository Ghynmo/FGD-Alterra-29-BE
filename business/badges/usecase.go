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

func (uc *BadgeUseCase) GetBadgesByPointController(ctx context.Context, point int) ([]Domain, error) {
	badge, err := uc.Repo.GetBadgesByPoint(ctx, point)
	if err != nil {
		return []Domain{}, err
	}

	return badge, nil
}
