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

func (uc *ThreadReportUseCase) GetReportsByCategoryController(ctx context.Context, category string) ([]Domain, error) {
	report, err := uc.Repo.GetReportsByCategory(ctx, category)
	if err != nil {
		return []Domain{}, err
	}

	return report, nil
}

func (uc *ThreadReportUseCase) GetThreadReports(ctx context.Context) ([]Domain, error) {
	report, err := uc.Repo.GetThreadReports(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return report, nil
}

func (uc *ThreadReportUseCase) CreateReportThread(ctx context.Context, domain Domain) (Domain, error) {
	report, err := uc.Repo.CreateReportThread(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return report, nil
}

func (uc *ThreadReportUseCase) GetReports(ctx context.Context) ([]Domain, error) {
	report, err := uc.Repo.GetReports(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return report, nil
}

func (uc *ThreadReportUseCase) DeleteThreadReport(ctx context.Context, id int) (Domain, error) {
	report, err := uc.Repo.DeleteThreadReport(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return report, nil
}
