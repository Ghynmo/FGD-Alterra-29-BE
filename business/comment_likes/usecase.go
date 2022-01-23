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

func (uc *CommentLikeUseCase) LikeController(ctx context.Context, domain Domain, id int) (Domain, error) {
	state, _ := uc.Repo.GetLikeState(ctx, domain, id)

	if state.Liker_id == 0 {
		comments, err := uc.Repo.NewLike(ctx, domain, id)
		if err != nil {
			return Domain{}, err
		}
		return comments, nil
	}

	if state.Liker_id != 0 && !state.State {
		comments, err := uc.Repo.Like(ctx, domain, id)
		if err != nil {
			return Domain{}, err
		}
		return comments, nil
	}

	if state.Liker_id != 0 && state.State {
		comments, err := uc.Repo.Unlike(ctx, domain, id)
		if err != nil {
			return Domain{}, err
		}
		return comments, nil
	}

	return Domain{}, nil
}
