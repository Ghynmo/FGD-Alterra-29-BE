package request

import "fgd-alterra-29/business/reputations"

type AddReputation struct {
	Reputation string `form:"reputation" json:"reputation"`
	LikePoint  int    `form:"likepoint" json:"likepoint"`
}

func (ar *AddReputation) ToDomain() reputations.Domain {
	return reputations.Domain{
		Reputation: ar.Reputation,
		LikePoints: ar.LikePoint,
	}
}
