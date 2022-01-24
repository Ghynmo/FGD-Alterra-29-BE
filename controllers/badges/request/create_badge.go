package request

import "fgd-alterra-29/business/badges"

type CreateBadge struct {
	Badge            string `form:"badge_name" json:"badge_name"`
	BadgeUrl         string `form:"badge_url" json:"badge_url"`
	RequirementPoint int    `form:"minimum_point" json:"minimum_point"`
	Description      string `form:"description" json:"description"`
}

func (cb *CreateBadge) ToDomain() badges.Domain {
	return badges.Domain{
		Badge:            cb.Badge,
		BadgeURL:         cb.BadgeUrl,
		RequirementPoint: cb.RequirementPoint,
		Description:      cb.Description,
	}
}
