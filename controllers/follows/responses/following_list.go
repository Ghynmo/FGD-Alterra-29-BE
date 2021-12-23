package responses

import "fgd-alterra-29/business/follows"

type FollowingList struct {
	Photo         string
	FollowingName string
	Reputation    string
}

func ToFollowingList(Domain follows.Domain) FollowingList {
	return FollowingList{
		Photo:         Domain.Photo,
		FollowingName: Domain.FollowingName,
		Reputation:    Domain.Reputation,
	}
}

func ToListFollowingList(u []follows.Domain) []FollowingList {
	var Domains []FollowingList

	for _, val := range u {
		Domains = append(Domains, ToFollowingList(val))
	}
	return Domains
}
