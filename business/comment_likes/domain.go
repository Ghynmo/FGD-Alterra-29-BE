package commentlikes

import (
	"context"
	"time"
)

type Domain struct {
	Comment_id int
	Liker_id   int
	Liked_at   time.Time
}

type UseCase interface {
	LikeController(ctx context.Context, domain Domain) (Domain, error)
	UnlikeController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	Like(ctx context.Context, domain Domain) (Domain, error)
	Unlike(ctx context.Context, domain Domain) (Domain, error)
}
