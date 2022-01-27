package responses

import (
	commentreport "fgd-alterra-29/business/comment_report"
)

type CommentReport struct {
	ID      int    `json:"report_id"`
	Comment string `json:"comment"`
	Message string `json:"message"`
	Case    string `json:"case"`
}

func ToReports(Domain commentreport.Domain) CommentReport {
	return CommentReport{
		ID:      Domain.ID,
		Comment: Domain.Comment,
		Message: Domain.Message,
		Case:    Domain.Case,
	}
}

func ToListReports(u []commentreport.Domain) []CommentReport {
	var Domains []CommentReport

	for _, val := range u {
		Domains = append(Domains, ToReports(val))
	}
	return Domains
}
