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

func (uc *CommentUseCase) GetPostsByCommentController(ctx context.Context, comment string) ([]Domain, error) {
	post, err := uc.Repo.GetPostsByComment(ctx, comment)
	if err != nil {
		return []Domain{}, err
	}

	return post, nil
}

func (uc *CommentUseCase) GetCommentProfile(ctx context.Context, id int) ([]Domain, error) {
	post, err := uc.Repo.GetCommentProfile(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return post, nil
}

func (uc *CommentUseCase) GetPostQuantity(ctx context.Context) (Domain, error) {
	post, err := uc.Repo.GetPostQuantity(ctx)
	if err != nil {
		return Domain{}, err
	}

	return post, nil
}

func (uc *CommentUseCase) GetPosts(ctx context.Context) ([]Domain, error) {
	post, err := uc.Repo.GetPosts(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return post, nil
}

func (uc *CommentUseCase) DeletePost(ctx context.Context, id int) (Domain, error) {
	post, err := uc.Repo.DeletePost(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return post, nil
}
