package follows

import (
	"context"
	"fgd-alterra-29/business/follows"
	"time"

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

	result := DB.Conn.Table("follows").Select("follower_id as Follower_id, photo_url as Photo, name as FollowerName, reputation").
		Where("user_id = (?) AND users.status = ?", id, "active").Joins("join users on follows.follower_id = users.id").
		Joins("join reputations on users.reputation_id = reputations.id").
		Order("followed_at desc").Find(&Follow)

	if result.Error != nil {
		return []follows.Domain{}, result.Error
	}

	return ToListDomain(Follow), nil
}

func (DB *MysqlFollowRepository) GetFollowing(ctx context.Context, id int) ([]follows.Domain, error) {
	var Follow []Follows

	result := DB.Conn.Table("follows").Select("user_id as User_id, photo_url as Photo, name as FollowingName, reputation").
		Where("follower_id = (?) AND users.status = ?", id, "active").Joins("join users on follows.user_id = users.id").
		Joins("join reputations on users.reputation_id = reputations.id").
		Order("followed_at desc").Find(&Follow)

	if result.Error != nil {
		return []follows.Domain{}, result.Error
	}

	return ToListDomain(Follow), nil
}

func (DB *MysqlFollowRepository) NewFollow(ctx context.Context, domain follows.Domain) (follows.Domain, error) {

	data := Follows{
		User_id:     domain.User_id,
		Follower_id: domain.Follower_id,
		Followed_at: time.Now(),
	}

	result := DB.Conn.Model(&data).Create(&data)

	if result.Error != nil {
		return follows.Domain{}, result.Error
	}

	return follows.Domain{}, nil
}

func (DB *MysqlFollowRepository) Follows(ctx context.Context, domain follows.Domain) (follows.Domain, error) {
	var follow Follows
	result := DB.Conn.Model(&follow).Where("user_id = ? AND follower_id = ?", domain.User_id, domain.Follower_id).
		Updates(Follows{State: true, Followed_at: time.Now()})

	if result.Error != nil {
		return follows.Domain{}, result.Error
	}

	return follow.ToDomain(), nil
}

func (DB *MysqlFollowRepository) Unfollow(ctx context.Context, domain follows.Domain) (follows.Domain, error) {
	var follow Follows

	result := DB.Conn.Model(&follow).Where("user_id = ? AND follower_id = ?", domain.User_id, domain.Follower_id).
		Update("state", false)

	if result.Error != nil {
		return follows.Domain{}, result.Error
	}

	return follow.ToDomain(), nil
}

func (DB *MysqlFollowRepository) GetFollowState(ctx context.Context, domain follows.Domain) (follows.Domain, error) {
	var follow Follows

	result := DB.Conn.Where("user_id = ? AND follower_id = ?", domain.User_id, domain.Follower_id).Find(&follow)

	if result.Error != nil {
		return follows.Domain{}, result.Error
	}

	return follow.ToDomain(), nil
}
