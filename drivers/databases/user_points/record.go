package userpoints

import userpoints "fgd-alterra-29/business/user_points"

type UserPoints struct {
	User_id     int `gorm:"unique"`
	ThreadPoint int `gorm:"default:0"`
	PostPoint   int `gorm:"default:0"`
}

func ToDomain(UserPoints UserPoints) userpoints.Domain {
	return userpoints.Domain{
		User_id:     UserPoints.User_id,
		ThreadPoint: UserPoints.ThreadPoint,
		PostPoint:   UserPoints.PostPoint,
	}
}
