package threads

import (
	"context"
	"fgd-alterra-29/business/comments"
	threadreport "fgd-alterra-29/business/thread_report"
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
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
	Category      string
	Comment       string
	Q_Comment     int
	RecentReplier string
	Q_Thread      int
	Name          string
	Photo         string
}

type UseCase interface {
	GetThreadsByTitleController(ctx context.Context, title string) ([]Domain, error)
	GetProfileThreads(ctx context.Context, id int) ([]Domain, error)
	GetThreadQuantity(ctx context.Context) (Domain, error)
	GetThreads(ctx context.Context) ([]Domain, error)
	DeleteThread(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetThreadsByTitle(ctx context.Context, title string) ([]Domain, error)
	GetProfileThreads(ctx context.Context, id int) ([]Domain, error)
	GetThreadQuantity(ctx context.Context) (Domain, error)
	GetThreads(ctx context.Context) ([]Domain, error)
	DeleteThread(ctx context.Context, id int) (Domain, error)
}
