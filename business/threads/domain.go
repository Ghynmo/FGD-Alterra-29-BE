package threads

import (
	"context"
	"fgd-alterra-29/business/comments"
	threadfollows "fgd-alterra-29/business/thread_follows"
	threadlikes "fgd-alterra-29/business/thread_likes"
	threadsaves "fgd-alterra-29/business/thread_saves"
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
	Followers     []threadfollows.Domain
	Saves         []threadsaves.Domain
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
	GetHomepageThreads(ctx context.Context, id int) ([]Domain, error)
	GetRecommendationThreads(ctx context.Context, id int) ([]Domain, error)
	GetHotThreads(ctx context.Context) ([]Domain, error)
	GetSearch(ctx context.Context, threadname string) ([]Domain, error)
}

type Repository interface {
	GetProfileThreads(ctx context.Context, id int) ([]Domain, error)
	GetHomepageThreads(ctx context.Context, id int) ([]Domain, error)
	GetRecommendationThreads(ctx context.Context, id int) ([]Domain, error)
	GetHotThreads(ctx context.Context) ([]Domain, error)
	GetSearch(ctx context.Context, threadname string) ([]Domain, error)
}
