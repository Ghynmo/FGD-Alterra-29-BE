package users

import (
	"context"
	commentlikes "fgd-alterra-29/business/comment_likes"
	"fgd-alterra-29/business/comments"
	"fgd-alterra-29/business/follows"
	threadfollows "fgd-alterra-29/business/thread_follows"
	threadlikes "fgd-alterra-29/business/thread_likes"
	threadsaves "fgd-alterra-29/business/thread_saves"
	"fgd-alterra-29/business/threads"
	userbadges "fgd-alterra-29/business/user_badges"
	"time"
)

type Domain struct {
	ID            int
	Role_id       int
	Reputation_id int
	Name          string
	Email         string
	Phone         string
	Address       string
	Header_url    string
	Photo_url     string
	Bio           string
	UserBadges    []userbadges.Domain
	Threads       []threads.Domain
	Following     []follows.Domain
	Followers     []follows.Domain
	Comments      []comments.Domain
	CommentLikes  []commentlikes.Domain
	ThreadLikes   []threadlikes.Domain
	ThreadFollows []threadfollows.Domain
	ThreadSaves   []threadsaves.Domain
	Q_Followers   int
	Q_Following   int
	Q_Post        int
	Q_Thread      int
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
}

type UseCase interface {
	GetUsersController(ctx context.Context) ([]Domain, error)
	GetProfileController(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetUsers(ctx context.Context) ([]Domain, error)
	GetProfile(ctx context.Context, id int) (Domain, error)
}
