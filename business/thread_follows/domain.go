package threadfollows

import "time"

type Domain struct {
	Thread_id   int
	User_id     int
	Followed_at time.Time
	Unfollow_at time.Time
}
