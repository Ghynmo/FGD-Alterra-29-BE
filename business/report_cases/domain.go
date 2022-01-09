package reportcases

import (
	"context"
	commentreport "fgd-alterra-29/business/comment_report"
	threadreport "fgd-alterra-29/business/thread_report"
)

type Domain struct {
	ID            int
	Case          string
	Description   string
	ThreadReport  []threadreport.Domain
	CommentReport []commentreport.Domain
}

type UseCase interface {
	GetReportForm(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	GetReportForm(ctx context.Context) ([]Domain, error)
}
