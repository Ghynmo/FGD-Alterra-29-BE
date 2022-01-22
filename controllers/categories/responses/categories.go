package responses

import "fgd-alterra-29/business/categories"

type Categories struct {
	Category string `json:"category"`
}

func ToCategories(c categories.Domain) Categories {
	return Categories{
		Category: c.Category,
	}
}

func ToListCategories(u []categories.Domain) []Categories {
	var Domains []Categories

	for _, val := range u {
		Domains = append(Domains, ToCategories(val))
	}
	return Domains
}
