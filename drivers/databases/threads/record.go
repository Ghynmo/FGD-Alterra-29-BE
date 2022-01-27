package threads

import (
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/drivers/databases/comments"
	threadfollows "fgd-alterra-29/drivers/databases/thread_follows"
	threadlikes "fgd-alterra-29/drivers/databases/thread_likes"
	threadreport "fgd-alterra-29/drivers/databases/thread_report"
	threadsaves "fgd-alterra-29/drivers/databases/thread_saves"
	threadshares "fgd-alterra-29/drivers/databases/thread_shares"
	"time"
)

type Threads struct {
	ID            int    `gorm:"primaryKey"`
	User_id       int    `gorm:"not null"`
	Category_id   int    `gorm:"not null"`
	Title         string `gorm:"not null"`
	Content       string `gorm:"not null"`
	Thumbnail_url string
	Active        bool                          `gorm:"default:true"`
	Report        []threadreport.ThreadReport   `gorm:"foreignKey:Thread_id"`
	Likes         []threadlikes.ThreadLikes     `gorm:"foreignKey:Thread_id"`
	Followers     []threadfollows.ThreadFollows `gorm:"foreignKey:Thread_id"`
	Saves         []threadsaves.ThreadSaves     `gorm:"foreignKey:Thread_id"`
	Shares        []threadshares.ThreadShares   `gorm:"foreignKey:Thread_id"`
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
	Name          string              `gorm:"-:migration;->"`
	Comments      []comments.Comments `gorm:"foreignKey:Thread_id"`
	Category      string              `gorm:"-:migration;->"`
	RecentReplier string              `gorm:"-:migration;->"`
	Comment       string              `gorm:"-:migration;->"`
	Q_Comment     int                 `gorm:"-:migration;->"`
	Q_Thread      int                 `gorm:"-:migration;->"`
	Photo         string              `gorm:"-:migration;->"`
	Q_Like        int                 `gorm:"-:migration;->"`
}

func (Thread *Threads) ToDomain() threads.Domain {
	return threads.Domain{
		ID:            Thread.ID,
		User_id:       Thread.User_id,
		Category_id:   Thread.Category_id,
		Title:         Thread.Title,
		Content:       Thread.Content,
		Thumbnail_url: Thread.Thumbnail_url,
		Active:        Thread.Active,
		Comments:      comments.ToListDomain(Thread.Comments),
		Created_at:    Thread.Created_at,
		Likes:         threadlikes.ToListDomain(Thread.Likes),
		Followers:     threadfollows.ToListDomain(Thread.Followers),
		Saves:         threadsaves.ToListDomain(Thread.Saves),
		Shares:        threadshares.ToListDomain(Thread.Shares),
		// Created_at:    Thread.Created_at,
		// Updated_at:    Thread.Updated_at,
		// Deleted_at:    Thread.Deleted_at,
		Name:          Thread.Name,
		Updated_at:    Thread.Updated_at,
		Deleted_at:    Thread.Deleted_at,
		Category:      Thread.Category,
		RecentReplier: Thread.RecentReplier,
		Comment:       Thread.Comment,
		Q_Comment:     Thread.Q_Comment,
		Q_Thread:      Thread.Q_Thread,
		Photo:         Thread.Photo,
		Q_Like:        Thread.Q_Like,
	}
}

func ToListDomain(u []Threads) []threads.Domain {
	var Domains []threads.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
