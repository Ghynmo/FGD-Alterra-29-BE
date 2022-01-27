package responses

import "fgd-alterra-29/business/follows"

type FollowingList struct {
	User_ID       int    `json:"user_id"`
	Photo         string `json:"photo"`
	FollowingName string `json:"name"`
	Reputation    string `json:"reputation"`
	FollowedByMe  bool   `json:"followed_by_me"`
}

func ToFollowingList(Domain follows.Domain) FollowingList {
	return FollowingList{
		User_ID:       Domain.User_id,
		Photo:         Domain.Photo,
		FollowingName: Domain.FollowingName,
		Reputation:    Domain.Reputation,
		FollowedByMe:  Domain.FollowedByMe,
	}
}

func ToListFollowingList(u []follows.Domain) []FollowingList {
	var Domains []FollowingList

	for _, val := range u {
		Domains = append(Domains, ToFollowingList(val))
	}
	return Domains
}
