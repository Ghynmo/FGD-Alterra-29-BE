package recommendation

import "fgd-alterra-29/business/comments"

type CommentRecommendation struct {
	Name    string
	Comment string
}

func ToCommentRecommendation(Domain comments.Domain) CommentRecommendation {
	return CommentRecommendation{
		Name:    Domain.Name,
		Comment: Domain.Comment,
	}
}

func ToListCommentRecommendation(u []comments.Domain) []CommentRecommendation {
	var Domains []CommentRecommendation

	for _, val := range u {
		Domains = append(Domains, ToCommentRecommendation(val))
	}
	return Domains
}
