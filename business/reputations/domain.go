package reputations

import "context"

type Domain struct {
	ID         int
	Reputation string
	LikePoints int
}

type UseCase interface {
	CreateReputationController(ctx context.Context, domain Domain) (Domain, error)
	GetReputationByUser(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	CreateReputation(ctx context.Context, domain Domain) (Domain, error)
	GetReputationByUser(ctx context.Context, id int) (Domain, error)
}
