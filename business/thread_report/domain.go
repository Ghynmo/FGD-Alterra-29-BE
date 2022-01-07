package threadreport

import (
	"context"
	"time"
)

type Domain struct {
	ID             int
	Thread_id      int
	User_id        int
	ReportGroup_id int
	Message        string
	Created_at     time.Time
	Updated_at     time.Time
	Deleted_at     time.Time
	Q_Cat          int
	CategoryReport string
	Thread         string
}

type UseCase interface {
	GetReportsByCategoryController(ctx context.Context, category string) ([]Domain, error)
	GetThreadReports(ctx context.Context) ([]Domain, error)
	CreateReportThread(ctx context.Context, domain Domain) (Domain, error)
	GetReports(ctx context.Context) ([]Domain, error)
	DeleteThreadReport(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetReportsByCategory(ctx context.Context, category string) ([]Domain, error)
	GetThreadReports(ctx context.Context) ([]Domain, error)
	CreateReportThread(ctx context.Context, domain Domain) (Domain, error)
	GetReports(ctx context.Context) ([]Domain, error)
	DeleteThreadReport(ctx context.Context, id int) (Domain, error)
}
