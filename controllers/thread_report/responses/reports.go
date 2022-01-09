package responses

import (
	threadreport "fgd-alterra-29/business/thread_report"
)

type ThreadReport struct {
	ID      int    `json:"report_id"`
	Thread  string `json:"title_thread"`
	Message string `json:"message"`
	Case    string `json:"report_case"`
}

func ToReports(Domain threadreport.Domain) ThreadReport {
	return ThreadReport{
		ID:      Domain.ID,
		Thread:  Domain.Thread,
		Message: Domain.Message,
		Case:    Domain.Case,
	}
}

func ToListReports(u []threadreport.Domain) []ThreadReport {
	var Domains []ThreadReport

	for _, val := range u {
		Domains = append(Domains, ToReports(val))
	}
	return Domains
}
