package profile

import "fgd-alterra-29/business/badges"

type ProfileBadges struct {
	Badge    string
	BadgeUrl string
}

func ToProfileBadges(Domain badges.Domain) ProfileBadges {
	return ProfileBadges{
		Badge:    Domain.Badge,
		BadgeUrl: Domain.BadgeURL,
	}
}

func ToListProfileBadges(u []badges.Domain) []ProfileBadges {
	var Domains []ProfileBadges

	for _, val := range u {
		Domains = append(Domains, ToProfileBadges(val))
	}
	return Domains
}
