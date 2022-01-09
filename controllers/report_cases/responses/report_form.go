package responses

import reportcases "fgd-alterra-29/business/report_cases"

type ReportForm struct {
	ID   int    `json:"reportcases_id"`
	Case string `json:"case"`
}

func ToReportForm(Domain reportcases.Domain) ReportForm {
	return ReportForm{
		ID:   Domain.ID,
		Case: Domain.Case,
	}
}

func ToListReportForm(u []reportcases.Domain) []ReportForm {
	var Domains []ReportForm

	for _, val := range u {
		Domains = append(Domains, ToReportForm(val))
	}
	return Domains
}
