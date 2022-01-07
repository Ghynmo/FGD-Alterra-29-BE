package responses

import "fgd-alterra-29/business/catreportthreads"

type ReportForm struct {
	ID             int    `json:"reportcat_id"`
	CategoryReport string `json:"report_category"`
}

func ToReportForm(Domain catreportthreads.Domain) ReportForm {
	return ReportForm{
		ID:             Domain.ID,
		CategoryReport: Domain.CategoryReport,
	}
}

func ToListReportForm(u []catreportthreads.Domain) []ReportForm {
	var Domains []ReportForm

	for _, val := range u {
		Domains = append(Domains, ToReportForm(val))
	}
	return Domains
}
