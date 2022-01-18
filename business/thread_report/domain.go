package threadreport

import (
	"context"
	"time"
)

type Domain struct {
	ID            int
	Thread_id     int
	Reporter_id   int
	ReportCase_id int
	Message       string
	State         string
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
	Q_Case        int
	Case          string
	Thread        string
}

type UseCase interface {
	SearchReportsByCategoryController(ctx context.Context, category string) ([]Domain, error)
	GetThreadReportStat(ctx context.Context) ([]Domain, error)
	CreateReportThread(ctx context.Context, domain Domain) (Domain, error)
	AdminGetReports(ctx context.Context) ([]Domain, error)
	SolvedThreadReport(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	SearchReportsByCategory(ctx context.Context, category string) ([]Domain, error)
	GetThreadReportStat(ctx context.Context) ([]Domain, error)
	CreateReportThread(ctx context.Context, domain Domain) (Domain, error)
	AdminGetReports(ctx context.Context) ([]Domain, error)
	SolvedThreadReport(ctx context.Context, id int) (Domain, error)
}
