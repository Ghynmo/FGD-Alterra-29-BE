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
	FollowingName string
	Reputation    string
}

type UseCase interface {
	GetFollowers(ctx context.Context, id int) ([]Domain, error)
	GetFollowing(ctx context.Context, id int) ([]Domain, error)
	FollowsController(ctx context.Context, domain Domain) (Domain, error)
	UnfollowController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	GetFollowers(ctx context.Context, id int) ([]Domain, error)
	GetFollowing(ctx context.Context, id int) ([]Domain, error)
	Follows(ctx context.Context, domain Domain) (Domain, error)
	Unfollow(ctx context.Context, domain Domain) (Domain, error)
}
