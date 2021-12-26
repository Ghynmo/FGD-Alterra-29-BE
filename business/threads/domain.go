package threads

import (
	"context"
	"fgd-alterra-29/business/comments"
	threadlikes "fgd-alterra-29/business/thread_likes"
	"time"
)

type Domain struct {
	ID            int
	User_id       int
	Category_id   int
	Title         string
	Content       string
	Thumbnail_url string
	Comments      []comments.Domain
	Likes         []threadlikes.Domain
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
	Name          string
	Category      string
	RecentReplier string
	Comment       string
	Q_Comment     int
	Q_Like        int
}

type UseCase interface {
	GetProfileThreads(ctx context.Context, id int) ([]Domain, error)
	GetRecommendationThreads(ctx context.Context, id int) ([]Domain, error)
}

type Repository interface {
	GetProfileThreads(ctx context.Context, id int) ([]Domain, error)
	GetRecommendationThreads(ctx context.Context, id int) ([]Domain, error)
}
