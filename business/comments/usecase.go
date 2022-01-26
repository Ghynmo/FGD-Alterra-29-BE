package comments

import (
	"context"
	up "fgd-alterra-29/business/user_points"
	"time"
)

type CommentUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
	UserPointRepo  up.Repository
}

func NewCommentUseCase(repo Repository, timeout time.Duration, up up.Repository) UseCase {
	return &CommentUseCase{
		Repo:           repo,
		contextTimeout: timeout,
		UserPointRepo:  up,
	}
}

func (uc *CommentUseCase) GetPostsByCommentController(ctx context.Context, comment string) ([]Domain, error) {
	post, err := uc.Repo.GetPostsByComment(ctx, comment)
	if err != nil {
		return []Domain{}, err

	}

	return post, nil
}

func (uc *CommentUseCase) GetCommentByThreadController(ctx context.Context, thread_id int, my_id int) ([]Domain, error) {
	comments, err := uc.Repo.GetCommentByThread(ctx, thread_id, my_id)
	if err != nil {
		return []Domain{}, err
	}

	return comments, nil
}

func (uc *CommentUseCase) GetCommentReply(ctx context.Context, id int, reply_of int) ([]Domain, error) {
	comments, err := uc.Repo.GetCommentReply(ctx, id, reply_of)
	if err != nil {
		return []Domain{}, err
	}

	return comments, nil
}

func (uc *CommentUseCase) GetCommentProfileController(ctx context.Context, id int) ([]Domain, error) {
	post, err := uc.Repo.GetCommentProfile(ctx, id)
	if err != nil {
		return []Domain{}, err
	}
	return post, nil
}

func (uc *CommentUseCase) GetPostQuantityController(ctx context.Context) (Domain, error) {
	post, err := uc.Repo.GetPostQuantity(ctx)
	if err != nil {
		return Domain{}, err
	}
	return post, nil
}

func (uc *CommentUseCase) CreateCommentController(ctx context.Context, domain Domain, id int) (Domain, error) {
	comments, user_id, err := uc.Repo.CreateComment(ctx, domain, id)
	uc.UserPointRepo.AddReputationPoint(ctx, 1, user_id)
	if err != nil {
		return Domain{}, err
	}

	return comments, nil
}

func (uc *CommentUseCase) GetPostsController(ctx context.Context) ([]Domain, error) {
	post, err := uc.Repo.GetPosts(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return post, nil
}

func (uc *CommentUseCase) UnactivatingPostController(ctx context.Context, id int) (Domain, error) {
	post, err := uc.Repo.UnactivatingPost(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return post, nil
}

func (uc *CommentUseCase) ActivatingPostController(ctx context.Context, id int) (Domain, error) {
	post, err := uc.Repo.ActivatingPost(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return post, nil
}
