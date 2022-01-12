package reputations

import "context"

type Domain struct {
	ID         int
	Reputation string
}

type UseCase interface {
	CreateReputationController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	CreateReputation(ctx context.Context, domain Domain) (Domain, error)
}
