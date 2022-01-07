package request

import threadreport "fgd-alterra-29/business/thread_report"

type CreateReport struct {
	Reporter_id    int    `form:"reporter_id"`
	Thread_id      int    `form:"thread_id"`
	CategoryReport int    `form:"category_report_id"`
	Message        string `form:"message"`
}

func (cr *CreateReport) ToDomain() threadreport.Domain {
	return threadreport.Domain{
		User_id:        cr.Reporter_id,
		Thread_id:      cr.Thread_id,
		ReportGroup_id: cr.CategoryReport,
		Message:        cr.Message,
	}
}
