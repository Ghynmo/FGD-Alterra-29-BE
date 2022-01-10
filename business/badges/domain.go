package badges

import userbadges "fgd-alterra-29/business/user_badges"

type Domain struct {
	ID          int
	Badge       string
	IconUrl     string
	Description string
	UserBadges  []userbadges.Domain
}
