package catreportthreads

import (
	"context"
	threadreport "fgd-alterra-29/business/thread_report"
)

type Domain struct {
	ID             int
	CategoryReport string
	Description    string
	ThreadReport   []threadreport.Domain
}

type UseCase interface {
	GetReportForm(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	GetReportForm(ctx context.Context) ([]Domain, error)
}
