package threadshares

import (
	threadshares "fgd-alterra-29/business/thread_shares"
)

type ThreadShares struct {
	Thread_id int `gorm:"not null"`
	User_id   int `gorm:"not null"`
}

func (TS *ThreadShares) ToDomain() threadshares.Domain {
	return threadshares.Domain{
		Thread_id: TS.Thread_id,
		User_id:   TS.User_id,
	}
}

func ToListDomain(u []ThreadShares) []threadshares.Domain {
	var Domains []threadshares.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
