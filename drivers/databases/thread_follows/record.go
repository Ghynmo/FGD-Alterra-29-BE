package threadfollows

import (
	threadfollows "fgd-alterra-29/business/thread_follows"
	"time"
)

type ThreadFollows struct {
	Thread_id     int `gorm:"not null"`
	User_id       int `gorm:"not null"`
	Followed_at   time.Time
	Unfollowed_at time.Time
}

func (TF *ThreadFollows) ToDomain() threadfollows.Domain {
	return threadfollows.Domain{
		Thread_id: TF.Thread_id,
		User_id:   TF.User_id,
	}
}

func ToListDomain(u []ThreadFollows) []threadfollows.Domain {
	var Domains []threadfollows.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
