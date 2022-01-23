package request

import "fgd-alterra-29/business/threads"

type CreateThread struct {
	Title       string `form:"title"`
	Category_id int    `form:"category_id"`
	Content     string `form:"content"`
}

func (ct *CreateThread) ToDomain() threads.Domain {
	return threads.Domain{
		Title:       ct.Title,
		Category_id: ct.Category_id,
		Content:     ct.Content,
	}
}
