package request

import (
	threadlikes "fgd-alterra-29/business/thread_likes"
)

type Like struct {
	Thread_id int `form:"thread_id" json:"thread_id"`
}

func (like *Like) ToDomain() threadlikes.Domain {
	return threadlikes.Domain{
		Thread_id: like.Thread_id,
	}
}
