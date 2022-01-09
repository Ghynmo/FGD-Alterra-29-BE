package request

import commentreport "fgd-alterra-29/business/comment_report"

type CreateReport struct {
	Reporter_id   int    `form:"reporter_id"`
	Comment_id    int    `form:"comment_id"`
	ReportCase_id int    `form:"report_case_id"`
	Message       string `form:"message"`
}

func (cr *CreateReport) ToDomain() commentreport.Domain {
	return commentreport.Domain{
		Reporter_id:   cr.Reporter_id,
		Comment_id:    cr.Comment_id,
		ReportCase_id: cr.ReportCase_id,
		Message:       cr.Message,
	}
}
