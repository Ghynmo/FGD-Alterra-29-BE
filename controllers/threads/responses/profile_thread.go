package responses

import "fgd-alterra-29/business/threads"

type ProfileThread struct {
	Title         string `json:"title"`
	Comment       string `json:"comment"`
	RecentReplier string `json:"recent_replier"`
}

func ToProfileThread(Domain threads.Domain) ProfileThread {
	return ProfileThread{
		Title:         Domain.Title,
		Comment:       Domain.Comment,
		RecentReplier: Domain.RecentReplier,
	}
}

func ToListProfileThread(u []threads.Domain) []ProfileThread {
	var Domains []ProfileThread

	for _, val := range u {
		Domains = append(Domains, ToProfileThread(val))
	}
	return Domains
}
