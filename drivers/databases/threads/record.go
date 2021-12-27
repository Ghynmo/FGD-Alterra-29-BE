package threads

import (
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/drivers/databases/comments"
	threadreport "fgd-alterra-29/drivers/databases/thread_report"
)

type Threads struct {
	ID            int `gorm:"primaryKey"`
	User_id       int
	Category_id   int
	Title         string
	Content       string
	Thumbnail_url string
	Comments      []comments.Comments         `gorm:"foreignKey:Thread_id"`
	Report        []threadreport.ThreadReport `gorm:"foreignKey:Thread_id"`
	// Created_at    time.Time
	// Updated_at    time.Time
	// Deleted_at    time.Time
	Category      string `gorm:"-:migration;->"`
	Comment       string `gorm:"-:migration;->"`
	Q_Comment     int    `gorm:"-:migration;->"`
	RecentReplier string `gorm:"-:migration;->"`
	Q_Thread      int    `gorm:"-:migration;->"`
}

func (Thread *Threads) ToDomain() threads.Domain {
	return threads.Domain{
		ID:            Thread.ID,
		User_id:       Thread.User_id,
		Category_id:   Thread.Category_id,
		Title:         Thread.Title,
		Content:       Thread.Content,
		Thumbnail_url: Thread.Thumbnail_url,
		Comments:      comments.ToListDomain(Thread.Comments),
		// Created_at:    Thread.Created_at,
		// Updated_at:    Thread.Updated_at,
		// Deleted_at:    Thread.Deleted_at,
		Category:      Thread.Category,
		Comment:       Thread.Comment,
		Q_Comment:     Thread.Q_Comment,
		RecentReplier: Thread.RecentReplier,
		Q_Thread:      Thread.Q_Thread,
	}
}

func ToListDomain(u []Threads) []threads.Domain {
	var Domains []threads.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
