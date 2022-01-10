package badges

import (
	"context"
	userbadges "fgd-alterra-29/business/user_badges"
)

type Domain struct {
	ID               int
	Badge            string
	BadgeURL         string
	Description      string
	RequirementPoint int
	UserBadges       []userbadges.Domain
}

type UseCase interface {
	GetBadgesByPointController(ctx context.Context, point int) ([]Domain, error)
}

type Repository interface {
	GetBadgesByPoint(ctx context.Context, point int) ([]Domain, error)
}
