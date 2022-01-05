package responses

import (
	"fgd-alterra-29/business/badges"
	"fgd-alterra-29/business/categories"
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/business/users"
	"fgd-alterra-29/controllers/users/responses/profile"
)

type Profile struct {
	ID             int
	Role_id        int
	Reputation_id  int
	Name           string
	Header_url     string
	Photo_url      string
	Bio            string
	Q_Following    int
	Q_Followers    int
	Q_Post         int
	Q_Thread       int
	UserBadgesTrue []profile.ProfileBadges    `json:"badge_list"`
	ActiveCategory []profile.ActiveOnCategory `json:"active_on_category"`
	ThreadProfile  []profile.ThreadProfile    `json:"thread_on_profile"`
}

func ToProfile(domain users.Domain, badge []badges.Domain, catthread []categories.Domain, thread []threads.Domain) Profile {
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
