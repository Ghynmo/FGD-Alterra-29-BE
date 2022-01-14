package threads

import (
	"context"
	"fgd-alterra-29/business/comments"
	threadfollows "fgd-alterra-29/business/thread_follows"
	threadlikes "fgd-alterra-29/business/thread_likes"
	threadreport "fgd-alterra-29/business/thread_report"
	threadsaves "fgd-alterra-29/business/thread_saves"
	threadshares "fgd-alterra-29/business/thread_shares"
	"time"
)

type Domain struct {
	ID            int
	User_id       int
	Category_id   int
	Title         string
	Content       string
	Thumbnail_url string
	Active        bool
	Comments      []comments.Domain
	Report        []threadreport.Domain
	Likes         []threadlikes.Domain
	Followers     []threadfollows.Domain
	Saves         []threadsaves.Domain
	Shares        []threadshares.Domain
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
	Name          string
	Category      string
	RecentReplier string
	Comment       string
	Q_Comment     int
	Q_Thread      int
	Photo         string
	Q_Like        int
}

type UseCase interface {
	GetThreadsByTitleController(ctx context.Context, title string) ([]Domain, error)
	GetProfileThreads(ctx context.Context, id int) ([]Domain, error)
	GetThreadQuantity(ctx context.Context) (Domain, error)
	GetThreads(ctx context.Context) ([]Domain, error)
	GetThreadByID(ctx context.Context, id int) (Domain, error)
	DeleteThread(ctx context.Context, id int) (Domain, error)
	ActivateThread(ctx context.Context, id int) (Domain, error)
	GetHomepageThreads(ctx context.Context, id int) ([]Domain, error)
	GetRecommendationThreads(ctx context.Context, id int) ([]Domain, error)
	GetHotThreads(ctx context.Context) ([]Domain, error)
	GetSearch(ctx context.Context, threadname string) ([]Domain, error)
	GetSideNewsThreads(ctx context.Context) ([]Domain, error)
	CreateThread(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	GetThreadsByTitle(ctx context.Context, title string) ([]Domain, error)
	GetProfileThreads(ctx context.Context, id int) ([]Domain, error)
	GetThreadQuantity(ctx context.Context) (Domain, error)
	GetThreads(ctx context.Context) ([]Domain, error)
	GetThreadByID(ctx context.Context, id int) (Domain, error)
	DeleteThread(ctx context.Context, id int) (Domain, error)
	ActivateThread(ctx context.Context, id int) (Domain, error)
	GetHomepageThreads(ctx context.Context, id int) ([]Domain, error)
	GetRecommendationThreads(ctx context.Context, id int) ([]Domain, error)
	GetHotThreads(ctx context.Context) ([]Domain, error)
	GetSearch(ctx context.Context, threadname string) ([]Domain, error)
	GetSideNewsThreads(ctx context.Context) ([]Domain, error)
	CreateThread(ctx context.Context, domain Domain) (Domain, error)
}
