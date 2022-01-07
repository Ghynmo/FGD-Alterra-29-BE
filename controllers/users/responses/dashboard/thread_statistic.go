package dashboard

import threadreport "fgd-alterra-29/business/thread_report"

type ThreadReportStat struct {
	ID             int    `json:"category_report_id"`
	CategoryReport string `json:"category_report"`
	Q_Cat          int    `json:"report_total"`
}

func ToThreadReportStat(Domain threadreport.Domain) ThreadReportStat {
	return ThreadReportStat{
		ID:             Domain.ID,
		CategoryReport: Domain.CategoryReport,
		Q_Cat:          Domain.Q_Cat,
	}
}

func ToListThreadReportStat(u []threadreport.Domain) []ThreadReportStat {
	var Domains []ThreadReportStat

	for _, val := range u {
		Domains = append(Domains, ToThreadReportStat(val))
	}
	return Domains
}
