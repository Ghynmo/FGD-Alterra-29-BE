package userpoints

import "context"

type Domain struct {
	User_id     int
	ThreadPoint int
	PostPoint   int
}

type UseCase interface {
	AddThreadPointController(ctx context.Context, id int) (Domain, error)
	AddPostPointController(ctx context.Context, id int) (Domain, error)
	AddReputationPointController(ctx context.Context, multiple int, id int) (Domain, error)
}

type Repository interface {
	AddThreadPoint(ctx context.Context, id int) (Domain, error)
	AddPostPoint(ctx context.Context, id int) (Domain, error)
	AddReputationPoint(ctx context.Context, multiple int, id int) (Domain, error)
}
