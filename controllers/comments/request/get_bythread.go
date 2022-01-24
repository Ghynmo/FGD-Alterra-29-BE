package request

type GetByThread struct {
	Thread_id int `form:"thread_id" json:"thread_id"`
}
