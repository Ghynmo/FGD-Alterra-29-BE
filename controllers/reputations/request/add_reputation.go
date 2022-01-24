package request

import "fgd-alterra-29/business/reputations"

type AddReputation struct {
	Reputation string `form:"reputation" json:"reputation"`
}

func (ar *AddReputation) ToDomain() reputations.Domain {
	return reputations.Domain{
		Reputation: ar.Reputation,
	}
}
