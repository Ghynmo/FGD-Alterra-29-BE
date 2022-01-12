package categories

import (
	"context"
	"time"
)

type CategoryUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewCategoryUseCase(repo Repository, timeout time.Duration) UseCase {
	return &CategoryUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *CategoryUseCase) GetUserActiveInCategory(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetUserActiveInCategory(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *CategoryUseCase) CreateCategoriesController(ctx context.Context, domain Domain) (Domain, error) {
	thread, err := uc.Repo.CreateCategories(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return thread, nil
}
