package commentlikes

import (
	commentlikes "fgd-alterra-29/business/comment_likes"
	"time"
)

type CommentLikes struct {
	Comment_id int
	Liker_id   int
	Liked_at   time.Time
	State      bool `gorm:"default:true"`
}

func (Cl *CommentLikes) ToDomain() commentlikes.Domain {
	return commentlikes.Domain{
		Comment_id: Cl.Comment_id,
		Liker_id:   Cl.Liker_id,
		Liked_at:   Cl.Liked_at,
		State:      Cl.State,
	}
}

func ToListDomain(u []CommentLikes) []commentlikes.Domain {
	var Domains []commentlikes.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
