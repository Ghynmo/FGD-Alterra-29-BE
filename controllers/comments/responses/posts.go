package responses

import (
	"fgd-alterra-29/business/comments"
)

type Posts struct {
	ID     int    `json:"post_id"`
	Name   string `json:"name"`
	Photo  string `json:"photo_url"`
	Post   string `json:"post"`
	Date   string `json:"created_at"`
	Active bool   `json:"active"`
}

func ToPosts(Domain comments.Domain) Posts {
	return Posts{
		ID:     Domain.ID,
		Name:   Domain.Name,
		Photo:  Domain.Photo,
		Post:   Domain.Comment,
		Date:   Domain.Created_at.String(),
		Active: Domain.Active,
	}
}

func ToListPosts(u []comments.Domain) []Posts {
	var Domains []Posts

	for _, val := range u {
		Domains = append(Domains, ToPosts(val))
	}
	return Domains
}
