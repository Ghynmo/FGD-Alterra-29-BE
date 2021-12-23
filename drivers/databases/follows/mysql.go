package follows

import (
	"context"
	"fgd-alterra-29/business/follows"
	"fmt"

	"gorm.io/gorm"
)

type MysqlFollowRepository struct {
	Conn *gorm.DB
}

func NewMysqlFollowRepository(conn *gorm.DB) follows.Repository {
	return &MysqlFollowRepository{
		Conn: conn,
	}
}

func (DB *MysqlFollowRepository) GetFollowers(ctx context.Context, id int) ([]follows.Domain, error) {
	var Follow []Follows

	result := DB.Conn.Table("follows").Select("photo_url as Photo, name as FollowerName, reputation").Where("user_id = (?)", id).
		Joins("join users on follows.follower_id = users.id").Joins("join reputations on users.reputation_id = reputations.id").
		Find(&Follow)

	if result.Error != nil {
		return []follows.Domain{}, result.Error
	}

	fmt.Println("Followers", Follow)
	return ToListDomain(Follow), nil
}

func (DB *MysqlFollowRepository) GetFollowing(ctx context.Context, id int) ([]follows.Domain, error) {
	var Follow []Follows

	result := DB.Conn.Table("follows").Select("photo_url as Photo, name as FollowingName, reputation").Where("follower_id = (?)", id).
		Joins("join users on follows.user_id = users.id").Joins("join reputations on users.reputation_id = reputations.id").
		Find(&Follow)

	if result.Error != nil {
		return []follows.Domain{}, result.Error
	}

	fmt.Println("Followers", Follow)
	return ToListDomain(Follow), nil
}
