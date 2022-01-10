package profile

import userbadges "fgd-alterra-29/business/user_badges"

type ProfileBadges struct {
	Badge    string `json:"badge"`
	Icon_url string `json:"icon"`
}

func ToProfileBadges(Domain userbadges.Domain) ProfileBadges {
	return ProfileBadges{
		Badge:    Domain.Badge,
		Icon_url: Domain.Icon_url,
	}
}

func ToListProfileBadges(u []userbadges.Domain) []ProfileBadges {
	var Domains []ProfileBadges

	for _, val := range u {
		Domains = append(Domains, ToProfileBadges(val))
	}
	return Domains
}
