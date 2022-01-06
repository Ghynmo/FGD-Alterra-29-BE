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
	comments, err := uc.Repo.GetCommentProfile(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return comments, nil
}

func (uc *CommentUseCase) GetCommentByThread(ctx context.Context, id int) ([]Domain, error) {
	comments, err := uc.Repo.GetCommentByThread(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return comments, nil
}

func (uc *CommentUseCase) CreateCommentController(ctx context.Context, domain Domain) (Domain, error) {
	comments, err := uc.Repo.CreateComment(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return comments, nil
}
