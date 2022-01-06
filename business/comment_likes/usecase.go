package commentlikes

import (
	"context"
	"time"
)

type CommentLikeUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
}

func NewCommentLikeUseCase(repo Repository, timeout time.Duration) UseCase {
	return &CommentLikeUseCase{
		Repo:           repo,
		contextTimeout: timeout,
	}
}

func (uc *CommentLikeUseCase) LikeController(ctx context.Context, domain Domain) (Domain, error) {
	comments, err := uc.Repo.Like(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return comments, nil
}

func (uc *CommentLikeUseCase) UnlikeController(ctx context.Context, domain Domain) (Domain, error) {
	comments, err := uc.Repo.Unlike(ctx, domain)
	if err != nil {
		return Domain{}, err
	}

	return comments, nil
}
