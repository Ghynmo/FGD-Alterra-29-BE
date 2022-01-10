package responses

import (
	"fgd-alterra-29/business/categories"
	"fgd-alterra-29/business/threads"
	userbadges "fgd-alterra-29/business/user_badges"
	"fgd-alterra-29/business/users"
	"fgd-alterra-29/controllers/users/responses/profile"
)

type Profile struct {
	ID             int                        `json:"user_id"`
	Role_id        int                        `json:"role_id"`
	Reputation_id  int                        `json:"reputation_id"`
	Name           string                     `json:"name"`
	Header_url     string                     `json:"header_url"`
	Photo_url      string                     `json:"photo_url"`
	Bio            string                     `json:"bio"`
	Q_Following    int                        `json:"following_total"`
	Q_Followers    int                        `json:"followers_total"`
	Q_Post         int                        `json:"post_total"`
	Q_Thread       int                        `json:"thread_total"`
	UserBadgesTrue []profile.ProfileBadges    `json:"badge_list"`
	ActiveCategory []profile.ActiveOnCategory `json:"active_on_category"`
	ThreadProfile  []profile.ThreadProfile    `json:"thread_on_profile"`
}

func ToProfile(domain users.Domain, badge []userbadges.Domain, catthread []categories.Domain, thread []threads.Domain) Profile {
	return Profile{
		ID:             domain.ID,
		Role_id:        domain.Role_id,
		Reputation_id:  domain.Reputation_id,
		Name:           domain.Name,
		Header_url:     domain.Header_url,
		Photo_url:      domain.Photo_url,
		Bio:            domain.Bio,
		Q_Following:    domain.Q_Following,
		Q_Followers:    domain.Q_Followers,
		Q_Post:         domain.Q_Post,
		Q_Thread:       domain.Q_Thread,
		UserBadgesTrue: profile.ToListProfileBadges(badge),
		ActiveCategory: profile.ToListActiveOnC(catthread),
		ThreadProfile:  profile.ToListThreadProfile(thread),
	}
}
