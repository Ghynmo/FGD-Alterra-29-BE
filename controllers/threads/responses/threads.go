package responses

import "fgd-alterra-29/business/threads"

type Threads struct {
	ID     int    `json:"thread_id"`
	Name   string `json:"name"`
	Photo  string `json:"photo_url"`
	Thread string `json:"thread_title"`
	Date   string `json:"created_at"`
}

func ToThread(Domain threads.Domain) Threads {
	return Threads{
		ID:     Domain.ID,
		Name:   Domain.Name,
		Photo:  Domain.Photo,
		Thread: Domain.Title,
		Date:   Domain.Created_at.String(),
	}
}

func ToListThread(u []threads.Domain) []Threads {
	var Domains []Threads

	for _, val := range u {
		Domains = append(Domains, ToThread(val))
	}
	return Domains
}
