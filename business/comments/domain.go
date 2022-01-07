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
	Photo      string
}

type UseCase interface {
	GetPostsByCommentController(ctx context.Context, comment string) ([]Domain, error)
	GetCommentProfile(ctx context.Context, id int) ([]Domain, error)
	GetPostQuantity(ctx context.Context) (Domain, error)
	GetPosts(ctx context.Context) ([]Domain, error)
	DeletePost(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetPostsByComment(ctx context.Context, comment string) ([]Domain, error)
	GetCommentProfile(ctx context.Context, id int) ([]Domain, error)
	GetPostQuantity(ctx context.Context) (Domain, error)
	GetPosts(ctx context.Context) ([]Domain, error)
	DeletePost(ctx context.Context, id int) (Domain, error)
}
