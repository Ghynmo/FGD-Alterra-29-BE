package profile

import "fgd-alterra-29/business/threads"

type ThreadProfile struct {
	Title         string
	RecentComment string
	RecentReplier string
}

func ToThreadProfile(Domain threads.Domain) ThreadProfile {
	return ThreadProfile{
		Title:         Domain.Title,
		RecentComment: Domain.Comment,
		RecentReplier: Domain.RecentReplier,
	}
}

func ToListThreadProfile(u []threads.Domain) []ThreadProfile {
	var Domains []ThreadProfile

	for _, val := range u {
		Domains = append(Domains, ToThreadProfile(val))
	}
	return Domains
}
