package threadlikes

import (
	"context"
	"time"
)

type Domain struct {
	Thread_id int
	User_id   int
	Liked_at  time.Time
	State     bool
}

type UseCase interface {
	LikeController(ctx context.Context, domain Domain, id int) (Domain, error)
}

type Repository interface {
	NewLike(ctx context.Context, domain Domain, id int) (Domain, int, error)
	Like(ctx context.Context, domain Domain, id int) (Domain, int, error)
	Unlike(ctx context.Context, domain Domain, id int) (Domain, int, error)
	GetLikeState(ctx context.Context, domain Domain, id int) (Domain, error)
}
