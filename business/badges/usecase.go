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

func (uc *BadgeUseCase) CreateBadgeController(ctx context.Context, domain Domain) (Domain, error) {
	badge, err := uc.Repo.CreateBadge(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return badge, nil
}

//Still Maintain
// func (uc *BadgeUseCase) BadgeStatusController(ctx context.Context, domain Domain) (Domain, error) {
// 	badge, err := uc.Repo.CreateBadge(ctx, domain)
// 	if err != nil {
// 		return Domain{}, err
// 	}

// 	return badge, nil
// }
