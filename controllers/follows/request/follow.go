package request

import (
	"fgd-alterra-29/business/follows"
)

type Follow struct {
	User_id     int `form:"user_id"`
	Follower_id int `form:"follower_id"`
}

func (follow *Follow) ToDomain() follows.Domain {
	return follows.Domain{
		User_id:     follow.User_id,
		Follower_id: follow.Follower_id,
	}
}
