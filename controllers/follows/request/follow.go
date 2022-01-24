package request

import (
	"fgd-alterra-29/business/follows"
)

type Follow struct {
	User_id int `form:"user_id" json:"user_id"`
}

func (follow *Follow) ToDomain() follows.Domain {
	return follows.Domain{
		User_id: follow.User_id,
	}
}
