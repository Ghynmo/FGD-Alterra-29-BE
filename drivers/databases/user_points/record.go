package userpoints

type UserPoints struct {
	User_id     int `gorm:"unique"`
	ThreadPoint int `gorm:"default:0"`
	PostPoint   int `gorm:"default:0"`
}
