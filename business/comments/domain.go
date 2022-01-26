package comments

import (
	"context"
	commentlikes "fgd-alterra-29/business/comment_likes"
	"time"
)

type Domain struct {
	ID         int
	Thread_id  int
	User_id    int
	Comment    string
	Replies    []Domain
	ReplyOf    int
	Active     bool
	Likes      []commentlikes.Domain
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
	Name       string
	Photo_url  string
	Thread     string
	Q_Post     int
	Photo      string
	LikeState  bool
}

type UseCase interface {
	GetPostsByCommentController(ctx context.Context, comment string) ([]Domain, error)
	GetCommentByThreadController(ctx context.Context, thread_id int, my_id int) ([]Domain, error)
	GetCommentProfileController(ctx context.Context, id int) ([]Domain, error)
	GetPostQuantityController(ctx context.Context) (Domain, error)
	GetPostsController(ctx context.Context) ([]Domain, error)
	UnactivatingPostController(ctx context.Context, id int) (Domain, error)
	ActivatingPostController(ctx context.Context, id int) (Domain, error)
	GetCommentReply(ctx context.Context, id int, reply_of int) ([]Domain, error)
	CreateCommentController(ctx context.Context, domain Domain, id int) (Domain, error)
}

type Repository interface {
	GetPostsByComment(ctx context.Context, comment string) ([]Domain, error)
	GetCommentByThread(ctx context.Context, thread_id int, my_id int) ([]Domain, error)
	GetCommentProfile(ctx context.Context, id int) ([]Domain, error)
	GetPostQuantity(ctx context.Context) (Domain, error)
	GetPosts(ctx context.Context) ([]Domain, error)
	UnactivatingPost(ctx context.Context, id int) (Domain, error)
	ActivatingPost(ctx context.Context, id int) (Domain, error)
	GetCommentReply(ctx context.Context, id int, reply_of int) ([]Domain, error)
	CreateComment(ctx context.Context, domain Domain, id int) (Domain, int, error)
}
