package badges

import (
	"fgd-alterra-29/business/badges"
	userbadges "fgd-alterra-29/drivers/databases/user_badges"
)

type Badges struct {
	ID                int    `gorm:"primaryKey"`
	Badge             string `gorm:"not null"`
	BadgeURL          string
	Description       string
	RequirementThread int
	UserBadges        []userbadges.UserBadges `gorm:"foreignKey:Badge_id"`
	Category_id       int
}

func (Badge *Badges) ToDomain() badges.Domain {
	return badges.Domain{
		ID:                Badge.ID,
		Badge:             Badge.Badge,
		BadgeURL:          Badge.BadgeURL,
		Description:       Badge.Description,
		RequirementThread: Badge.RequirementThread,
		UserBadges:        userbadges.ToListDomain(Badge.UserBadges),
		Category_id:       Badge.Category_id,
	}
}

func ToListDomain(u []Badges) []badges.Domain {
	var Domains []badges.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
