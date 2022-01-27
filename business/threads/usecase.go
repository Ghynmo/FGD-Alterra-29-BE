package threads

import (
	"context"
	"fgd-alterra-29/business/badges"
	ub "fgd-alterra-29/business/user_badges"
	"time"
)

type ThreadUseCase struct {
	Repo           Repository
	contextTimeout time.Duration
	RepoBadges     badges.Repository
	RepoUserBadges ub.Repository
}

func NewThreadUseCase(repo Repository, timeout time.Duration, repobadge badges.Repository, ub ub.Repository) UseCase {
	return &ThreadUseCase{
		Repo:           repo,
		contextTimeout: timeout,
		RepoBadges:     repobadge,
		RepoUserBadges: ub,
	}
}

func (uc *ThreadUseCase) GetThreadsByTitleController(ctx context.Context, title string) ([]Domain, error) {
	thread, err := uc.Repo.GetThreadsByTitle(ctx, title)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetProfileThreads(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetProfileThreads(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetThreadQuantity(ctx context.Context) (Domain, error) {
	thread, err := uc.Repo.GetThreadQuantity(ctx)
	if err != nil {
		return Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) CreateThread(ctx context.Context, domain Domain, id int) (Domain, error) {
	thread, err := uc.Repo.CreateThread(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	qty, err := uc.Repo.GetThreadQtyByCategory(ctx, domain, id)
	if err != nil {
		return Domain{}, err
	}

	newbadge, err3 := uc.RepoBadges.GetBadgesIdByThread(ctx, qty.Q_Thread)
	if err3 != nil {
		return Domain{}, err3
	}

	checkBadgeID, err4 := uc.RepoUserBadges.CheckGetBadge(ctx, id, newbadge)
	if err4 != nil {
		return Domain{}, err3
	}
	if checkBadgeID.User_id == 0 {
		_, err4 := uc.RepoUserBadges.CreatenewRecord(ctx, id, newbadge)
		if err4 != nil {
			return Domain{}, err4
		}
		return thread, nil
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetHomepageThreads(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetHomepageThreads(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetThreads(ctx context.Context) ([]Domain, error) {
	thread, err := uc.Repo.GetThreads(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetRecommendationThreads(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetRecommendationThreads(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetThreadByID(ctx context.Context, id int) (Domain, error) {
	thread, err := uc.Repo.GetThreadByID(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) DeleteThread(ctx context.Context, id int) (Domain, error) {
	thread, err := uc.Repo.DeleteThread(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) ActivateThread(ctx context.Context, id int) (Domain, error) {
	thread, err := uc.Repo.ActivateThread(ctx, id)
	if err != nil {
		return Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetHotThreads(ctx context.Context) ([]Domain, error) {
	thread, err := uc.Repo.GetHotThreads(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetSearch(ctx context.Context, threadname string) ([]Domain, error) {
	thread, err := uc.Repo.GetSearch(ctx, threadname)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetThreadsByCategoryID(ctx context.Context, id int) ([]Domain, error) {
	thread, err := uc.Repo.GetThreadsByCategoryID(ctx, id)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}

func (uc *ThreadUseCase) GetSideNewsThreads(ctx context.Context) ([]Domain, error) {
	thread, err := uc.Repo.GetSideNewsThreads(ctx)
	if err != nil {
		return []Domain{}, err
	}

	return thread, nil
}
