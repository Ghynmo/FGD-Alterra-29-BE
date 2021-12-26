package threadlikes

import "time"

type Domain struct {
	Thread_id  int
	User_id    int
	Liked_at   time.Time
	Unliked_at time.Time
}
