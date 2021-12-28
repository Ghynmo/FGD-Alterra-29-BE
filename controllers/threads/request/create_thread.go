package request

import "fgd-alterra-29/business/threads"

type CreateThread struct {
	User_id     int    `form:"user_id"`
	Title       string `form:"title"`
	Category_id int    `form:"category_id"`
	Content     string `form:"content"`
}

func (ct *CreateThread) ToDomain() threads.Domain {
	return threads.Domain{
		User_id:     ct.User_id,
		Title:       ct.Title,
		Category_id: ct.Category_id,
		Content:     ct.Content,
	}
}
