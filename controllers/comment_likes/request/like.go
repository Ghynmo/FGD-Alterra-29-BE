package request

import commentlikes "fgd-alterra-29/business/comment_likes"

type Like struct {
	Comment_id int `form:"comment_id"`
	Liker_id   int `form:"liker_id"`
}

func (like *Like) ToDomain() commentlikes.Domain {
	return commentlikes.Domain{
		Comment_id: like.Comment_id,
		Liker_id:   like.Liker_id,
	}
}
