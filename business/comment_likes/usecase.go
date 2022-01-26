package commentlikes

import (
	"context"
	up "fgd-alterra-29/business/user_points"
	"time"
)

type CommentLikeUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
	UserPointRepo  up.Repository
}

func NewCommentLikeUseCase(repo Repository, timeout time.Duration, up up.Repository) UseCase {
	return &CommentLikeUseCase{
		Repo:           repo,
		contextTimeout: timeout,
		UserPointRepo:  up,
	}
}

func (uc *CommentLikeUseCase) LikeController(ctx context.Context, domain Domain, id int) (Domain, error) {
	state, errState := uc.Repo.GetLikeState(ctx, domain, id)
	if errState != nil {
		return Domain{}, errState
	}

	if state.Liker_id == 0 {
		comments, t_user_id, err := uc.Repo.NewLike(ctx, domain, id)
		uc.UserPointRepo.AddReputationPoint(ctx, 1, t_user_id)
		if err != nil {
			return Domain{}, err
		}
		return comments, nil
	}

	if state.Liker_id != 0 && !state.State {
		comments, t_user_id, err := uc.Repo.Like(ctx, domain, id)
		uc.UserPointRepo.AddReputationPoint(ctx, 1, t_user_id)
		if err != nil {
			return Domain{}, err
		}
		return comments, nil
	}

	if state.Liker_id != 0 && state.State {
		comments, t_user_id, err := uc.Repo.Unlike(ctx, domain, id)
		uc.UserPointRepo.AddReputationPoint(ctx, -1, t_user_id)
		if err != nil {
			return Domain{}, err
		}
		return comments, nil
	}

	return Domain{}, nil
}
