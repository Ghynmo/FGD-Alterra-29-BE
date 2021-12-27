package dashboard

import threadreport "fgd-alterra-29/business/thread_report"

type ThreadReportStat struct {
	CategoryReport string
	Q_Cat          int
}

func ToThreadReportStat(Domain threadreport.Domain) ThreadReportStat {
	return ThreadReportStat{
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
