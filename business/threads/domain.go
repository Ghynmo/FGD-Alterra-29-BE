package threads

import (
	"context"
	"fgd-alterra-29/business/comments"
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
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
	Category      string
	Q_Title       int
	Comment       string
	Q_Comment     int
	RecentReplier string
}

type UseCase interface {
	GetProfileThreads(ctx context.Context, id int) ([]Domain, error)
}

type Repository interface {
	GetProfileThreads(ctx context.Context, id int) ([]Domain, error)
}
