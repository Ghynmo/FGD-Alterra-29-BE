package threadreport

import "context"

type Domain struct {
	ID             int
	Thread_id      int
	User_id        int
	ReportGroup_id int
	Message        string
	Q_Cat          int
	CategoryReport string
}

type UseCase interface {
	GetThreadReports(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	GetThreadReports(ctx context.Context) ([]Domain, error)
}
