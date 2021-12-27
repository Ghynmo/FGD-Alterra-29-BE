package catreportthreads

import threadreport "fgd-alterra-29/business/thread_report"

type Domain struct {
	ID             int
	CategoryReport string
	Description    string
	ThreadReport   []threadreport.Domain
}
