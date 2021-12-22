package follows

import (
	"fgd-alterra-29/business/follows"
	"time"
)

type Follows struct {
	User_id       int `gorm:"not null"`
	Follower_id   int `gorm:"not null"`
	Followed_at   time.Time
	Unfollowed_at time.Time
}

func (Follow *Follows) ToDomain() follows.Domain {
	return follows.Domain{
		User_id:       Follow.User_id,
		Follower_id:   Follow.Follower_id,
		Followed_at:   Follow.Followed_at,
		Unfollowed_at: Follow.Unfollowed_at,
	}
}

func ToListDomain(u []Follows) []follows.Domain {
	var Domains []follows.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
