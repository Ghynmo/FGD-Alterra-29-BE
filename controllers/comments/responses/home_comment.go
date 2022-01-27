package responses

import (
	"fgd-alterra-29/business/comments"
)

type ReplyComment struct {
	ID      int    `json:"comment_id"`
	Name    string `json:"name"`
	Photo   string `json:"photo_url"`
	Comment string `json:"comment"`
}

func ToReplyComment(Domain comments.Domain) ReplyComment {
	return ReplyComment{
		ID:      Domain.ID,
		Name:    Domain.Name,
		Photo:   Domain.Photo_url,
		Comment: Domain.Comment,
	}
}

func ToListReplyComment(u []comments.Domain) []ReplyComment {
	var Domains []ReplyComment

	for _, val := range u {
		Domains = append(Domains, ToReplyComment(val))
	}
	return Domains
}
