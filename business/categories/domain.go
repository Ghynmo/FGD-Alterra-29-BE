package categories

import (
	"context"
	"fgd-alterra-29/business/threads"
)

type Domain struct {
	ID       int
	Category string
	IconUrl  string
	Threads  []threads.Domain
	Q_Title  int
}

type UseCase interface {
	GetUserActiveInCategory(ctx context.Context, id int) ([]Domain, error)
	CreateCategoriesController(ctx context.Context, domain Domain) (Domain, error)
}

type Repository interface {
	GetUserActiveInCategory(ctx context.Context, id int) ([]Domain, error)
	CreateCategories(ctx context.Context, domain Domain) (Domain, error)
}
