package request

import (
	threadlikes "fgd-alterra-29/business/thread_likes"
)

type Like struct {
	Thread_id int `form:"thread_id"`
	User_id   int `form:"user_id"`
}

func (like *Like) ToDomain() threadlikes.Domain {
	return threadlikes.Domain{
		Thread_id: like.Thread_id,
		User_id:   like.User_id,
	}
}
