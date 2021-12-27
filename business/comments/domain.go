package comments

import (
	"context"
	"time"
)

type Domain struct {
	ID         int
	Thread_id  int
	User_id    int
	Comment    string
	Replies    []Domain
	ReplyOf    int
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
	Name       string
	Thread     string
	Q_Post     int
}

type UseCase interface {
	GetCommentProfile(ctx context.Context, id int) ([]Domain, error)
	GetPostQuantity(ctx context.Context) (Domain, error)
}

type Repository interface {
	GetCommentProfile(ctx context.Context, id int) ([]Domain, error)
	GetPostQuantity(ctx context.Context) (Domain, error)
}
