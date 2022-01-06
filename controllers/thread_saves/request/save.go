package request

import (
	threadsaves "fgd-alterra-29/business/thread_saves"
)

type Save struct {
	Thread_id int `form:"thread_id"`
	User_id   int `form:"user_id"`
}

func (like *Save) ToDomain() threadsaves.Domain {
	return threadsaves.Domain{
		Thread_id: like.Thread_id,
		User_id:   like.User_id,
	}
}
