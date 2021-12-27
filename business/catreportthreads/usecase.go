package catreportthreads

import (
	"context"
	"time"
)

type CatReportThreadUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewCatReportThreadUseCase(repo Repository, timeout time.Duration) UseCase {
	return &CatReportThreadUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *CatReportThreadUseCase) GetReportForm(ctx context.Context) ([]Domain, error) {
	catreportthread, err := uc.Repo.GetReportForm(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return catreportthread, nil
}
