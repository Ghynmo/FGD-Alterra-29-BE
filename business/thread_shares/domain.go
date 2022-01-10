package threadshares

import "context"

type Domain struct {
	Thread_id int
	User_id   int
}

type UseCase interface {
	ThreadShareController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	ThreadShare(ctx context.Context, domain Domain) (Domain, error)
}
