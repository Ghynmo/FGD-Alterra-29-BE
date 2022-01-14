package responses

import apinews "fgd-alterra-29/business/api_news"

type Articles struct {
	Articles []SideNews `json:"article"`
}

func ToDomain(Domain apinews.Domain) Articles {
	return Articles{
		Articles: ToListSideNews(Domain.Articles),
	}
}

type SideNews struct {
	Title   string `json:"title"`
	NewsUrl string `json:"url"`
	ImgUrl  string `json:"urlToImage"`
}

func ToSideNews(Subdomain apinews.SubDomain) SideNews {
	return SideNews{
		Title:   Subdomain.Title,
		NewsUrl: Subdomain.NewsUrl,
		ImgUrl:  Subdomain.ImgUrl,
	}
}

func ToListSideNews(u []apinews.SubDomain) []SideNews {
	var Domains []SideNews

	for _, val := range u {
		Domains = append(Domains, ToSideNews(val))
	}
	return Domains
}
