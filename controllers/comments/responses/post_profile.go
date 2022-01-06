package responses

import (
	"fgd-alterra-29/business/comments"
)

type PostProfile struct {
	TargetName string `json:"comment_target"`
	Thread     string `json:"thread"`
	Comment    string `json:"comment"`
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
