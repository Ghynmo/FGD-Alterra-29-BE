package threadlikes

import (
	"context"
	"time"
)

type Domain struct {
	Thread_id  int
	User_id    int
	Liked_at   time.Time
	Unliked_at time.Time
}

type UseCase interface {
	LikeController(ctx context.Context, domain Domain) (Domain, error)
	UnlikeController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	Like(ctx context.Context, domain Domain) (Domain, error)
	Unlike(ctx context.Context, domain Domain) (Domain, error)
}
