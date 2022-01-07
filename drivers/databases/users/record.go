package users

import (
	"fgd-alterra-29/business/users"
	commentlikes "fgd-alterra-29/drivers/databases/comment_likes"
	"fgd-alterra-29/drivers/databases/comments"
	"fgd-alterra-29/drivers/databases/follows"
	threadfollows "fgd-alterra-29/drivers/databases/thread_follows"
	threadlikes "fgd-alterra-29/drivers/databases/thread_likes"
	threadsaves "fgd-alterra-29/drivers/databases/thread_saves"
	threadshares "fgd-alterra-29/drivers/databases/thread_shares"
	"fgd-alterra-29/drivers/databases/threads"
	userbadges "fgd-alterra-29/drivers/databases/user_badges"
)

type Users struct {
	ID            int `gorm:"primaryKey"`
	Role_id       int
	Reputation_id int
	Name          string `gorm:"not null"`
	Email         string
	Phone         string
	Address       string
	Header_url    string
	Photo_url     string
	Bio           string
	UserBadges    []userbadges.UserBadges       `gorm:"foreignKey:User_id"`
	Threads       []threads.Threads             `gorm:"foreignKey:User_id"`
	Following     []follows.Follows             `gorm:"foreignKey:User_id"`
	Followers     []follows.Follows             `gorm:"foreignKey:Follower_id"`
	Comments      []comments.Comments           `gorm:"foreignKey:User_id"`
	CommentLikes  []commentlikes.CommentLikes   `gorm:"foreignKey:Liker_id"`
	ThreadLikes   []threadlikes.ThreadLikes     `gorm:"foreignKey:User_id"`
	ThreadFollows []threadfollows.ThreadFollows `gorm:"foreignKey:User_id"`
	ThreadSaves   []threadsaves.ThreadSaves     `gorm:"foreignKey:User_id"`
	ThreadShares  []threadshares.ThreadShares   `gorm:"foreignKey:User_id"`
	Q_Following   int                           `gorm:"-:migration;->"`
	Q_Followers   int                           `gorm:"-:migration;->"`
	Q_Post        int                           `gorm:"-:migration;->"`
	Q_Thread      int                           `gorm:"-:migration;->"`
	Reputation    string                        `gorm:"-:migration;->"`
	// Created_at    time.Time
	// Updated_at    time.Time
	// Deleted_at    time.Time
}

func (user *Users) ToDomain() users.Domain {
	return users.Domain{
		ID:            user.ID,
		Role_id:       user.Role_id,
		Reputation_id: user.Reputation_id,
		Name:          user.Name,
		Email:         user.Email,
		Phone:         user.Phone,
		Address:       user.Address,
		Header_url:    user.Header_url,
		Photo_url:     user.Photo_url,
		Bio:           user.Bio,
		UserBadges:    userbadges.ToListDomain(user.UserBadges),
		Threads:       threads.ToListDomain(user.Threads),
		Following:     follows.ToListDomain(user.Following),
		Followers:     follows.ToListDomain(user.Followers),
		Comments:      comments.ToListDomain(user.Comments),
		CommentLikes:  commentlikes.ToListDomain(user.CommentLikes),
		ThreadLikes:   threadlikes.ToListDomain(user.ThreadLikes),
		ThreadFollows: threadfollows.ToListDomain(user.ThreadFollows),
		ThreadSaves:   threadsaves.ToListDomain(user.ThreadSaves),
		ThreadShares:  threadshares.ToListDomain(user.ThreadShares),
		Q_Following:   user.Q_Following,
		Q_Followers:   user.Q_Followers,
		Q_Post:        user.Q_Post,
		Q_Thread:      user.Q_Thread,
		// Created_at:    user.Created_at,
		// Updated_at:    user.Updated_at,
		// Deleted_at:    user.Deleted_at,
	}
}

func ToListDomain(u []Users) []users.Domain {
	var Domains []users.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
