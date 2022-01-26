package request

import "fgd-alterra-29/business/badges"

type CreateBadge struct {
	Badge             string `form:"badge_name" json:"badge_name"`
	Category_id       int    `form:"category_id" json:"category_id"`
	BadgeUrl          string `form:"badge_url" json:"badge_url"`
	RequirementThread int    `form:"thread_qty" json:"thread_qty"`
	Description       string `form:"description" json:"description"`
}

func (cb *CreateBadge) ToDomain() badges.Domain {
	return badges.Domain{
		Badge:             cb.Badge,
		Category_id:       cb.Category_id,
		BadgeURL:          cb.BadgeUrl,
		RequirementThread: cb.RequirementThread,
		Description:       cb.Description,
	}
}
