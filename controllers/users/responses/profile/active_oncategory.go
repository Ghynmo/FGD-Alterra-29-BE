package profile

import "fgd-alterra-29/business/categories"

type ActiveOnCategory struct {
	Category string
	Q_Title  int
}

func ToActiveOnC(Domain categories.Domain) ActiveOnCategory {
	return ActiveOnCategory{
		Category: Domain.Category,
		Q_Title:  Domain.Q_Title,
	}
}

func ToListActiveOnC(u []categories.Domain) []ActiveOnCategory {
	var Domains []ActiveOnCategory

	for _, val := range u {
		Domains = append(Domains, ToActiveOnC(val))
	}
	return Domains
}
