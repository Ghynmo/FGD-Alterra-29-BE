package threadsaves

import (
	"context"
	"time"
)

type Domain struct {
	Thread_id int
	User_id   int
	Saved_at  time.Time
}

type UseCase interface {
	SaveThreadController(ctx context.Context, domain Domain) (Domain, error)
	UnsaveThreadController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	SaveThread(ctx context.Context, domain Domain) (Domain, error)
	UnsaveThread(ctx context.Context, domain Domain) (Domain, error)
}
