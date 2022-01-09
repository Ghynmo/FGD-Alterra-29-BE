package dashboard

import threadreport "fgd-alterra-29/business/thread_report"

type ThreadReportStat struct {
	ID     int    `json:"category_report_id"`
	Case   string `json:"case"`
	Q_Case int    `json:"report_total"`
}

func ToThreadReportStat(Domain threadreport.Domain) ThreadReportStat {
	return ThreadReportStat{
		ID:     Domain.ID,
		Case:   Domain.Case,
		Q_Case: Domain.Q_Case,
	}
}

func ToListThreadReportStat(u []threadreport.Domain) []ThreadReportStat {
	var Domains []ThreadReportStat

	for _, val := range u {
		Domains = append(Domains, ToThreadReportStat(val))
	}
	return Domains
}
