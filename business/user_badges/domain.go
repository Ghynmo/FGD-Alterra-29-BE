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
}
