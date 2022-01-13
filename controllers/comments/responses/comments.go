package responses

import "fgd-alterra-29/business/comments"

type CommentRecommendation struct {
	ID        int    `json:"comment_id"`
	Name      string `json:"name"`
	Photo_url string `json:"photo_url"`
	Comment   string `json:"comment"`
	LikeState bool   `json:"likes_by_me"`
}

func ToCommentRecommendation(Domain comments.Domain) CommentRecommendation {
	return CommentRecommendation{
		ID:        Domain.ID,
		Name:      Domain.Name,
		Photo_url: Domain.Photo_url,
		Comment:   Domain.Comment,
		LikeState: Domain.LikeState,
	}
}

func ToListCommentRecommendation(u []comments.Domain) []CommentRecommendation {
	var Domains []CommentRecommendation

	for _, val := range u {
		Domains = append(Domains, ToCommentRecommendation(val))
	}
	return Domains
}
