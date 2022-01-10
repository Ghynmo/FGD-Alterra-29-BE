package badges

import (
	"fgd-alterra-29/business/badges"
	userbadges "fgd-alterra-29/drivers/databases/user_badges"
)

type Badges struct {
	ID          int    `gorm:"primaryKey"`
	Badge       string `gorm:"not null"`
	IconUrl     string
	Description string
	UserBadges  []userbadges.UserBadges `gorm:"foreignKey:Badge_id"`
}

func (Badge *Badges) ToDomain() badges.Domain {
	return badges.Domain{
		ID:          Badge.ID,
		Badge:       Badge.Badge,
		IconUrl:     Badge.IconUrl,
		Description: Badge.Description,
		UserBadges:  userbadges.ToListDomain(Badge.UserBadges),
	}
}

func ToListDomain(u []Badges) []badges.Domain {
	var Domains []badges.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
