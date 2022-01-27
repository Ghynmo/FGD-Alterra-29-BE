package threadsaves

import (
	"context"
	"time"
)

type Domain struct {
	Thread_id int
	User_id   int
	Saved_at  time.Time
	State     bool
}

type UseCase interface {
	SaveThreadController(ctx context.Context, domain Domain, id int) (Domain, error)
}

type Repository interface {
	NewSave(ctx context.Context, domain Domain, id int) (Domain, error)
	Save(ctx context.Context, domain Domain, id int) (Domain, error)
	Unsave(ctx context.Context, domain Domain, id int) (Domain, error)
	GetSaveState(ctx context.Context, domain Domain, id int) (Domain, error)
}
