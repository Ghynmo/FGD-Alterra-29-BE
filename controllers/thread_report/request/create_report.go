package request

import threadreport "fgd-alterra-29/business/thread_report"

type CreateReport struct {
	Thread_id     int    `form:"thread_id" json:"thread_id"`
	ReportCase_id int    `form:"report_case_id" json:"report_case_id"`
	Message       string `form:"message" json:"message"`
}

func (cr *CreateReport) ToDomain() threadreport.Domain {
	return threadreport.Domain{
		Thread_id:     cr.Thread_id,
		ReportCase_id: cr.ReportCase_id,
		Message:       cr.Message,
	}
}
