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
	Photo_url  string
	Thread     string
}

type UseCase interface {
	GetCommentProfile(ctx context.Context, id int) ([]Domain, error)
	GetCommentByThread(ctx context.Context, id int) ([]Domain, error)
	CreateCommentController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	GetCommentProfile(ctx context.Context, id int) ([]Domain, error)
	GetCommentByThread(ctx context.Context, id int) ([]Domain, error)
	CreateComment(ctx context.Context, domain Domain) (Domain, error)
}
