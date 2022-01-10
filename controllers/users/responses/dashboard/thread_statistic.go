package dashboard

import commentreport "fgd-alterra-29/business/comment_report"

type CommentReportStat struct {
	ID     int    `json:"report_case_id"`
	Case   string `json:"case"`
	Q_Case int    `json:"report_total"`
}

func ToCommentReportStat(Domain commentreport.Domain) CommentReportStat {
	return CommentReportStat{
		ID:     Domain.ID,
		Case:   Domain.Case,
		Q_Case: Domain.Q_Case,
	}
}

func ToListCommentReportStat(u []commentreport.Domain) []CommentReportStat {
	var Domains []CommentReportStat

	for _, val := range u {
		Domains = append(Domains, ToCommentReportStat(val))
	}
	return Domains
}
