package request

type GetFollows struct {
	Target_id int `form:"target_id"`
	My_id     int `form:"my_user_id"`
}
