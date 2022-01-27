package newsapi

import (
	apinews "fgd-alterra-29/business/api_news"
)

type ArticlesParent struct {
	Articles []SideNews `json:"articles"`
}

func (Ap *ArticlesParent) ToDomain() apinews.Domain {
	return apinews.Domain{
		Articles: ToListDomain(Ap.Articles),
	}
}
