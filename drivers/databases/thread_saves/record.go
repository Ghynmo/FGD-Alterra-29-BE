package threadsaves

import (
	threadsaves "fgd-alterra-29/business/thread_saves"
	"time"
)

type ThreadSaves struct {
	Thread_id int `gorm:"not null"`
	User_id   int `gorm:"not null"`
	Saved_at  time.Time
}

func (TS *ThreadSaves) ToDomain() threadsaves.Domain {
	return threadsaves.Domain{
		Thread_id: TS.Thread_id,
		User_id:   TS.User_id,
		Saved_at:  TS.Saved_at,
	}
}

func ToListDomain(u []ThreadSaves) []threadsaves.Domain {
	var Domains []threadsaves.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
