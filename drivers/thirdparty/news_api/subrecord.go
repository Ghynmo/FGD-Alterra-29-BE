package newsapi

import apinews "fgd-alterra-29/business/api_news"

type SideNews struct {
	Author       string `json:"author"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	NewsUrl      string `json:"url"`
	ImgUrl       string `json:"urlToImage"`
	Published_at string `json:"published_at"`
	Content      string `json:"content"`
}

func (News *SideNews) ToDomain() apinews.SubDomain {
	return apinews.SubDomain{
		Title:   News.Title,
		NewsUrl: News.NewsUrl,
		ImgUrl:  News.ImgUrl,
	}
}

func ToListDomain(u []SideNews) []apinews.SubDomain {
	var Domains []apinews.SubDomain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
