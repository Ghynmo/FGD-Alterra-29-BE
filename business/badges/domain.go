package badges

import (
	"context"
	userbadges "fgd-alterra-29/business/user_badges"
)

type Domain struct {
	ID                int
	Badge             string
	BadgeURL          string
	Description       string
	RequirementThread int
	Status            bool
	Category_id       int
	UserBadges        []userbadges.Domain
}

type UseCase interface {
	GetBadgesByUserController(ctx context.Context, id int) ([]Domain, error)
	CreateBadgeController(ctx context.Context, domain Domain) (Domain, error)
	// BadgeStatusController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	GetBadgesIdByThread(ctx context.Context, thread_qty int) (int, error)
	GetBadgesByUser(ctx context.Context, id int) ([]Domain, error)
	CreateBadge(ctx context.Context, domain Domain) (Domain, error)
	ActivateBadge(ctx context.Context, domain Domain) (Domain, error)
	UnactivateBadge(ctx context.Context, domain Domain) (Domain, error)
}
