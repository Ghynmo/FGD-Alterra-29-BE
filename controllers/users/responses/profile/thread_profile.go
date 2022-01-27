package profile

import "fgd-alterra-29/business/threads"

type ThreadProfile struct {
	ID            int    `json:"thread_id"`
	Title         string `json:"title"`
	RecentComment string `json:"recent_comment"`
	RecentReplier string `json:"recent_replier"`
}

func ToThreadProfile(Domain threads.Domain) ThreadProfile {
	return ThreadProfile{
		ID:            Domain.ID,
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
