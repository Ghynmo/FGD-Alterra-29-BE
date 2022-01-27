package threadshares

import (
	"context"
	"time"
)

type Domain struct {
	Thread_id int
	User_id   int
	Shared_at time.Time
}

type UseCase interface {
	ThreadShareController(ctx context.Context, domain Domain, id int) (Domain, error)
}

type Repository interface {
	ThreadShare(ctx context.Context, domain Domain, id int) (Domain, error)
}
