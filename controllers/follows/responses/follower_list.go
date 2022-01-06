package responses

import "fgd-alterra-29/business/follows"

type FollowerList struct {
	User_ID      int    `json:"user_id"`
	Photo        string `json:"photo"`
	FollowerName string `json:"name"`
	Reputation   string `json:"reputation"`
}

func ToFollowerList(Domain follows.Domain) FollowerList {
	return FollowerList{
		User_ID:      Domain.Follower_id,
		Photo:        Domain.Photo,
		FollowerName: Domain.FollowerName,
		Reputation:   Domain.Reputation,
	}
}

func ToListFollowerList(u []follows.Domain) []FollowerList {
	var Domains []FollowerList

	for _, val := range u {
		Domains = append(Domains, ToFollowerList(val))
	}
	return Domains
}
