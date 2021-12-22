package follows

import (
	"time"
)

type Domain struct {
	User_id       int
	Follower_id   int
	Followed_at   time.Time
	Unfollowed_at time.Time
}
