package comments

import (
	"context"
	"time"
)

type CommentUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewCommentUseCase(repo Repository, timeout time.Duration) UseCase {
	return &CommentUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *CommentUseCase) GetCommentProfile(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetCommentProfile(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}
