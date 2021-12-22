package userbadges

import userbadges "fgd-alterra-29/business/user_badges"

type UserBadges struct {
	User_id  int
	Badge_id int
	Badge    string `gorm:"-:migration;->"`
}

func (UserBadges *UserBadges) ToDomain() userbadges.Domain {
	return userbadges.Domain{
		User_id:  UserBadges.User_id,
		Badge_id: UserBadges.Badge_id,
		Badge:    UserBadges.Badge,
	}
}

func ToListDomain(u []UserBadges) []userbadges.Domain {
	var Domains []userbadges.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
