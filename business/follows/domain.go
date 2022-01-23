package follows

import (
	"context"
	"time"
)

type Domain struct {
	User_id       int
	Follower_id   int
	State         bool
	Followed_at   time.Time
	Unfollowed_at time.Time
	Photo         string
	FollowerName  string
	FollowingName string
	Reputation    string
	FollowedByMe  bool
}

type UseCase interface {
	GetFollowers(ctx context.Context, target_id int, my_id int) ([]Domain, error)
	GetFollowing(ctx context.Context, target_id int, my_id int) ([]Domain, error)
	FollowsController(ctx context.Context, domain Domain, my_id int) (Domain, error)
}

type Repository interface {
	GetFollowers(ctx context.Context, target_id int, my_id int) ([]Domain, error)
	GetFollowing(ctx context.Context, target_id int, my_id int) ([]Domain, error)
	Follows(ctx context.Context, domain Domain) (Domain, error)
	Unfollow(ctx context.Context, domain Domain) (Domain, error)
	NewFollow(ctx context.Context, domain Domain) (Domain, error)
	GetFollowState(ctx context.Context, domain Domain, my_id int) (Domain, error)
}
