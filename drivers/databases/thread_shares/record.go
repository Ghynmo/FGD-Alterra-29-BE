package threadshares

import (
	threadshares "fgd-alterra-29/business/thread_shares"
	"time"
)

type ThreadShares struct {
	Thread_id int `gorm:"not null"`
	User_id   int `gorm:"not null"`
	Shared_at time.Time
}

func (TS *ThreadShares) ToDomain() threadshares.Domain {
	return threadshares.Domain{
		Thread_id: TS.Thread_id,
		User_id:   TS.User_id,
		Shared_at: TS.Shared_at,
	}
}

func ToListDomain(u []ThreadShares) []threadshares.Domain {
	var Domains []threadshares.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
