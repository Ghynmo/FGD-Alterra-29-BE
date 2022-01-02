package home

import "fgd-alterra-29/business/comments"

type CommentRecommendation struct {
	Name      string
	Photo_url string
	Comment   string
}

func ToCommentRecommendation(Domain comments.Domain) CommentRecommendation {
	return CommentRecommendation{
		Name:      Domain.Name,
		Photo_url: Domain.Photo_url,
		Comment:   Domain.Comment,
	}
}

func ToListCommentRecommendation(u []comments.Domain) []CommentRecommendation {
	var Domains []CommentRecommendation

	for _, val := range u {
		Domains = append(Domains, ToCommentRecommendation(val))
	}
	return Domains
}
