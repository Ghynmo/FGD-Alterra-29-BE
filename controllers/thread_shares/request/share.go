package request

import (
	threadshares "fgd-alterra-29/business/thread_shares"
)

type Share struct {
	Thread_id int `form:"thread_id" json:"thread_id"`
}

func (like *Share) ToDomain() threadshares.Domain {
	return threadshares.Domain{
		Thread_id: like.Thread_id,
	}
}
