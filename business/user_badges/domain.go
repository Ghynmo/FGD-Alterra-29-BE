package userbadges

import "context"

type Domain struct {
	User_id  int
	Badge_id int
	Badge    string
	Icon_url string
}

type UseCase interface {
	GetUserBadge(ctx context.Context, id int) ([]Domain, error)
}

type Repository interface {
	GetUserBadge(ctx context.Context, id int) ([]Domain, error)
	CheckGetBadge(ctx context.Context, user_id int, badge_id int) (Domain, error)
	CreatenewRecord(ctx context.Context, user_id int, badge_id int) (Domain, error)
}
