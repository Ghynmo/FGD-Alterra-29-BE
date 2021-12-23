package responses

import "fgd-alterra-29/business/follows"

type FollowerList struct {
	Photo        string
	FollowerName string
	Reputation   string
}

func ToFollowerList(Domain follows.Domain) FollowerList {
	return FollowerList{
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
