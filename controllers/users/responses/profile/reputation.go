package profile

import (
	"fgd-alterra-29/business/reputations"
)

type Reputation struct {
	Reputation string `json:"reputation"`
}

func ToNewReputation(Domain reputations.Domain) Reputation {
	return Reputation{
		Reputation: Domain.Reputation,
	}
}
