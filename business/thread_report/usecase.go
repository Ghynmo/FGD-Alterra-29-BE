package threadreport

import (
	"context"
	"time"
)

type ThreadReportUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewThreadReportUseCase(repo Repository, timeout time.Duration) UseCase {
	return &ThreadReportUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *ThreadReportUseCase) GetThreadReports(ctx context.Context) ([]Domain, error) {
	report, err := uc.Repo.GetThreadReports(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return report, nil
}
