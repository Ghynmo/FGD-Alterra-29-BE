package responses

import (
	"fgd-alterra-29/business/threads"
)

type RecommendationThreads struct {
	ID        int    `json:"thread_id"`
	UserID    int    `json:"user_id"`
	Name      string `json:"thread_maker"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	Q_Like    int    `json:"likes_total"`
	Q_Comment int    `json:"comments_total"`
	// Comments  []home.CommentRecommendation `json:"comments"`
}

func ToHomeThreads(Domain threads.Domain) RecommendationThreads {
	return RecommendationThreads{
		ID:        Domain.ID,
		UserID:    Domain.User_id,
		Name:      Domain.Name,
		Title:     Domain.Title,
		Content:   Domain.Content,
		Q_Like:    Domain.Q_Like,
		Q_Comment: Domain.Q_Comment,
		// Comments:  home.ToListCommentRecommendation(Domain.Comments),
	}
}

func ToListHomeThreads(t []threads.Domain) []RecommendationThreads {
	var Domains []RecommendationThreads

	for _, val := range t {
		Domains = append(Domains, ToHomeThreads(val))
	}
	return Domains
}
