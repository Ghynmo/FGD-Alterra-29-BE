package commentreport

import (
	"context"
	"time"
)

type Domain struct {
	ID            int
	Comment_id    int
	Reporter_id   int
	ReportCase_id int
	Message       string
	Created_at    time.Time
	Q_Case        int
	Case          string
	Comment       string
}

type UseCase interface {
	SearchReportsByCategoryController(ctx context.Context, category string) ([]Domain, error)
	GetCommentReportStat(ctx context.Context) ([]Domain, error)
	CreateReportComment(ctx context.Context, domain Domain) (Domain, error)
	AdminGetReports(ctx context.Context) ([]Domain, error)
	DeleteCommentReport(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	SearchReportsByCategory(ctx context.Context, category string) ([]Domain, error)
	GetCommentReportStat(ctx context.Context) ([]Domain, error)
	CreateReportComment(ctx context.Context, domain Domain) (Domain, error)
	AdminGetReports(ctx context.Context) ([]Domain, error)
	DeleteCommentReport(ctx context.Context, id int) (Domain, error)
}
