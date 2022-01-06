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
}

type Repository interface {
	GetUserActiveInCategory(ctx context.Context, id int) ([]Domain, error)
}
