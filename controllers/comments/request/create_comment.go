package request

import "fgd-alterra-29/business/comments"

type CreateComment struct {
	Thread_id int    `form:"thread_id" json:"thread_id"`
	Reply_of  int    `form:"reply_of" json:"reply_of"`
	Comment   string `form:"comment" json:"comment"`
}

func (cc *CreateComment) ToDomain() comments.Domain {
	return comments.Domain{
		Thread_id: cc.Thread_id,
		ReplyOf:   cc.Reply_of,
		Comment:   cc.Comment,
	}
}
