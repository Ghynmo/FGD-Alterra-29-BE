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

func (uc *ThreadReportUseCase) SearchReportsByCategoryController(ctx context.Context, category string) ([]Domain, error) {
	report, err := uc.Repo.SearchReportsByCategory(ctx, category)
	if err != nil {
		return []Domain{}, err
	}

	return report, nil
}

func (uc *ThreadReportUseCase) GetThreadReportStat(ctx context.Context) ([]Domain, error) {
	report, err := uc.Repo.GetThreadReportStat(ctx)
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

func (uc *ThreadReportUseCase) AdminGetReports(ctx context.Context) ([]Domain, error) {
	report, err := uc.Repo.AdminGetReports(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return report, nil
}

func (uc *ThreadReportUseCase) SolvedThreadReport(ctx context.Context, id int) (Domain, error) {
	report, err := uc.Repo.SolvedThreadReport(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return report, nil
}
