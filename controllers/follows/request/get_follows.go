package request

type GetFollows struct {
	Target_id int `form:"target_id" json:"target_id"`
}
