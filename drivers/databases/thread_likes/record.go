package threadlikes

import (
	threadlikes "fgd-alterra-29/business/thread_likes"
	"time"
)

type ThreadLikes struct {
	Thread_id int
	User_id   int
	Liked_at  time.Time
	State     bool `gorm:"default:true"`
}

func (TL *ThreadLikes) ToDomain() threadlikes.Domain {
	return threadlikes.Domain{
		Thread_id: TL.Thread_id,
		User_id:   TL.User_id,
		State:     TL.State,
	}
}

func ToListDomain(u []ThreadLikes) []threadlikes.Domain {
	var Domains []threadlikes.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
