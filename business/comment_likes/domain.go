package commentlikes

import (
	"context"
	"time"
)

type Domain struct {
	Comment_id int
	Liker_id   int
	Liked_at   time.Time
	State      bool
}

type UseCase interface {
	LikeController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	NewLike(ctx context.Context, domain Domain) (Domain, error)
	Like(ctx context.Context, domain Domain) (Domain, error)
	Unlike(ctx context.Context, domain Domain) (Domain, error)
	GetLikeState(ctx context.Context, domain Domain) (Domain, error)
}