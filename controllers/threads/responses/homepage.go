package responses

import (
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/controllers/threads/responses/home"
)

type RecommendationThreads struct {
	Name      string
	Title     string
	Content   string
	Q_Like    int
	Q_Comment int
	Comments  []home.CommentRecommendation
}

func ToRecommendationThreads(Domain threads.Domain) RecommendationThreads {
	return RecommendationThreads{
		Name:      Domain.Name,
		Title:     Domain.Title,
		Content:   Domain.Content,
		Q_Like:    Domain.Q_Like,
		Q_Comment: Domain.Q_Comment,
		Comments:  home.ToListCommentRecommendation(Domain.Comments),
	}
}

func ToListRecommendationThreads(t []threads.Domain) []RecommendationThreads {
	var Domains []RecommendationThreads

	for _, val := range t {
		Domains = append(Domains, ToRecommendationThreads(val))
	}
	return Domains
}
