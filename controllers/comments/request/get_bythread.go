package request

type GetByThread struct {
	Thread_id int `form:"thread_id"`
	MyUser_id int `form:"myuser_id"`
}
