package responses

import (
	"fgd-alterra-29/business/comments"
)

type PostProfile struct {
	TargetName string
	Thread     string
	Comment    string
}

func ToPostProfile(Domain comments.Domain) PostProfile {
	return PostProfile{
		TargetName: Domain.Name,
		Thread:     Domain.Thread,
		Comment:    Domain.Comment,
	}
}

func ToListPostProfile(u []comments.Domain) []PostProfile {
	var Domains []PostProfile

	for _, val := range u {
		Domains = append(Domains, ToPostProfile(val))
	}
	return Domains
}
