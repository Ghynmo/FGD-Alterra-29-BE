package reputations

import "fgd-alterra-29/business/reputations"

type Reputations struct {
	ID         int    `gorm:"primaryKey"`
	Reputation string `gorm:"not null"`
}

func (Reputation *Reputations) ToDomain() reputations.Domain {
	return reputations.Domain{
		ID:         Reputation.ID,
		Reputation: Reputation.Reputation,
	}
}
