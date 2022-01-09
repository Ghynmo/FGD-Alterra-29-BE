package request

import threadreport "fgd-alterra-29/business/thread_report"

type CreateReport struct {
	Reporter_id   int    `form:"reporter_id"`
	Thread_id     int    `form:"thread_id"`
	ReportCase_id int    `form:"report_case_id"`
	Message       string `form:"message"`
}

func (cr *CreateReport) ToDomain() threadreport.Domain {
	return threadreport.Domain{
		Reporter_id:   cr.Reporter_id,
		Thread_id:     cr.Thread_id,
		ReportCase_id: cr.ReportCase_id,
		Message:       cr.Message,
	}
}
