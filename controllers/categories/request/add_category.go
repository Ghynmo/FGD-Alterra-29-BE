package request

import "fgd-alterra-29/business/categories"

type AddCategory struct {
	Category string `form:"category"`
}

func (ac *AddCategory) ToDomain() categories.Domain {
	return categories.Domain{
		Category: ac.Category,
	}
}
