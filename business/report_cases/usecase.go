package reportcases

import (
	"context"
	"time"
)

type ReportCaseUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewReportCaseUseCase(repo Repository, timeout time.Duration) UseCase {
	return &ReportCaseUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *ReportCaseUseCase) GetReportForm(ctx context.Context) ([]Domain, error) {
	catreportthread, err := uc.Repo.GetReportForm(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return catreportthread, nil
}

func (uc *ReportCaseUseCase) CreateCaseController(ctx context.Context, domain Domain) (Domain, error) {
	catreportthread, err := uc.Repo.CreateCase(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return catreportthread, nil
}
