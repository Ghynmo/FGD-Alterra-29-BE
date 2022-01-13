package follows

import (
	"fgd-alterra-29/business/follows"
	"time"
)

type Follows struct {
	User_id       int  `gorm:"not null"`
	Follower_id   int  `gorm:"not null"`
	State         bool `gorm:"default:true"`
	Followed_at   time.Time
	Unfollowed_at time.Time
	Photo         string `gorm:"-:migration;->"`
	FollowerName  string `gorm:"-:migration;->"`
	FollowingName string `gorm:"-:migration;->"`
	Reputation    string `gorm:"-:migration;->"`
}

func (Follow *Follows) ToDomain() follows.Domain {
	return follows.Domain{
		User_id:       Follow.User_id,
		Follower_id:   Follow.Follower_id,
		State:         Follow.State,
		Followed_at:   Follow.Followed_at,
		Unfollowed_at: Follow.Unfollowed_at,
		Photo:         Follow.Photo,
		FollowerName:  Follow.FollowerName,
		FollowingName: Follow.FollowingName,
		Reputation:    Follow.Reputation,
	}
}

func ToListDomain(u []Follows) []follows.Domain {
	var Domains []follows.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
