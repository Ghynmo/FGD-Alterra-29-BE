package commentreport

import (
	"context"
	"time"
)

type CommentReportUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewCommentReportUseCase(repo Repository, timeout time.Duration) UseCase {
	return &CommentReportUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *CommentReportUseCase) SearchReportsByCategoryController(ctx context.Context, category string) ([]Domain, error) {
	report, err := uc.Repo.SearchReportsByCategory(ctx, category)
	if err != nil {
		return []Domain{}, err
	}

	return report, nil
}

func (uc *CommentReportUseCase) GetCommentReportStat(ctx context.Context) ([]Domain, error) {
	report, err := uc.Repo.GetCommentReportStat(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return report, nil
}

func (uc *CommentReportUseCase) CreateReportComment(ctx context.Context, domain Domain) (Domain, error) {
	report, err := uc.Repo.CreateReportComment(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return report, nil
}

func (uc *CommentReportUseCase) AdminGetReports(ctx context.Context) ([]Domain, error) {
	report, err := uc.Repo.AdminGetReports(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return report, nil
}

func (uc *CommentReportUseCase) DeleteCommentReport(ctx context.Context, id int) (Domain, error) {
	report, err := uc.Repo.DeleteCommentReport(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return report, nil
}
