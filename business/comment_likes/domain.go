package commentlikes

import "context"

type Domain struct {
	Comment_id int
	Liker_id   int
}

type UseCase interface {
	LikeController(ctx context.Context, domain Domain) (Domain, error)
	UnlikeController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	Like(ctx context.Context, domain Domain) (Domain, error)
	Unlike(ctx context.Context, domain Domain) (Domain, error)
}
