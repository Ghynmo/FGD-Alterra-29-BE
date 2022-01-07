package comments

import (
	"fgd-alterra-29/business/comments"
	"time"
)

type Comments struct {
	ID         int `gorm:"primaryKey"`
	Thread_id  int
	User_id    int
	Comment    string
	Replies    []Comments `gorm:"foreignKey:ReplyOf"`
	ReplyOf    int
	Active     bool
	Created_at time.Time
	Updated_at time.Time
	Deleted_at time.Time
	Name       string `gorm:"-:migration;->"`
	Thread     string `gorm:"-:migration;->"`
	Q_Post     int    `gorm:"-:migration;->"`
	Photo      string `gorm:"-:migration;->"`
}

func (Comment *Comments) ToDomain() comments.Domain {
	return comments.Domain{
		ID:         Comment.ID,
		Thread_id:  Comment.Thread_id,
		User_id:    Comment.User_id,
		Comment:    Comment.Comment,
		Replies:    ToListDomain(Comment.Replies),
		ReplyOf:    Comment.ReplyOf,
		Active:     Comment.Active,
		Created_at: Comment.Created_at,
		Updated_at: Comment.Updated_at,
		Deleted_at: Comment.Deleted_at,
		Name:       Comment.Name,
		Thread:     Comment.Thread,
		Q_Post:     Comment.Q_Post,
		Photo:      Comment.Photo,
	}
}

func ToListDomain(u []Comments) []comments.Domain {
	var Domains []comments.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
