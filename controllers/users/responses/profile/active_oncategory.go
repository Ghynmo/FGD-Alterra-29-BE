package profile

import "fgd-alterra-29/business/categories"

type ActiveOnCategory struct {
	Category string `json:"category"`
	Q_Title  int    `json:"thread_quantity"`
	IconUrl  string `json:"icon"`
}

func ToActiveOnC(Domain categories.Domain) ActiveOnCategory {
	return ActiveOnCategory{
		Category: Domain.Category,
		Q_Title:  Domain.Q_Title,
		IconUrl:  Domain.IconUrl,
	}
}

func ToListActiveOnC(u []categories.Domain) []ActiveOnCategory {
	var Domains []ActiveOnCategory

	for _, val := range u {
		Domains = append(Domains, ToActiveOnC(val))
	}
	return Domains
}
