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
	GetThreadReports(ctx context.Context) ([]Domain, error)
	GetReports(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	GetThreadReports(ctx context.Context) ([]Domain, error)
	GetReports(ctx context.Context) ([]Domain, error)
}
