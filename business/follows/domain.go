package follows

import (
	"context"
	"time"
)

type Domain struct {
	User_id       int
	Follower_id   int
	Followed_at   time.Time
	Unfollowed_at time.Time
	Photo         string
	FollowerName  string
	Reputation    string
}

type UseCase interface {
	GetFollowers(ctx context.Context, id int) ([]Domain, error)
}

type Repository interface {
	GetFollowers(ctx context.Context, id int) ([]Domain, error)
}
