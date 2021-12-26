package threadlikes

import (
	threadlikes "fgd-alterra-29/business/thread_likes"
	"time"
)

type ThreadLikes struct {
	Thread_id  int `gorm:"not null"`
	User_id    int `gorm:"not null"`
	Liked_at   time.Time
	Unliked_at time.Time
}

func (TL *ThreadLikes) ToDomain() threadlikes.Domain {
	return threadlikes.Domain{
		Thread_id: TL.Thread_id,
		User_id:   TL.User_id,
	}
}

func ToListDomain(u []ThreadLikes) []threadlikes.Domain {
	var Domains []threadlikes.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
