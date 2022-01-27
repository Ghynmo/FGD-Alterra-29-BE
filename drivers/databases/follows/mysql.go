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

func (DB *MysqlFollowRepository) GetFollowers(ctx context.Context, target_id int, my_id int) ([]follows.Domain, error) {
	var Follow []Follows

	result := DB.Conn.Raw("SELECT follows.follower_id as Follower_id, photo_url as Photo, name as FollowerName, reputation, nestedfollow.state AS FollowedByMe FROM follows LEFT JOIN (SELECT user_id, follower_id, state FROM follows WHERE follower_id = ?) AS nestedfollow ON follows.follower_id = nestedfollow.user_id JOIN users ON follows.follower_id = users.id JOIN reputations ON users.reputation_id = reputations.id WHERE follows.user_id = ? AND follows.follower_id != ? ORDER BY follows.followed_at DESC",
		my_id, target_id, my_id).
		Scan(&Follow)

	if result.Error != nil {
		return []follows.Domain{}, result.Error
	}

	return ToListDomain(Follow), nil
}

func (DB *MysqlFollowRepository) GetFollowing(ctx context.Context, target_id int, my_id int) ([]follows.Domain, error) {
	var Follow []Follows

	result := DB.Conn.Raw("SELECT follows.user_id as User_id, photo_url as Photo, name as FollowingName, reputation, nestedfollow.state AS FollowedByMe FROM follows LEFT JOIN (SELECT user_id, follower_id, state FROM follows WHERE follower_id = ?) AS nestedfollow ON follows.user_id = nestedfollow.user_id JOIN users ON follows.user_id = users.id JOIN reputations ON users.reputation_id = reputations.id WHERE follows.follower_id = ? AND follows.user_id != ? ORDER BY follows.followed_at DESC",
		my_id, target_id, my_id).
		Scan(&Follow)

	if result.Error != nil {
		return []follows.Domain{}, result.Error
	}

	return ToListDomain(Follow), nil
}

func (DB *MysqlFollowRepository) NewFollow(ctx context.Context, domain follows.Domain, my_id int) (follows.Domain, error) {

	data := Follows{
		User_id:     domain.User_id,
		Follower_id: my_id,
		Followed_at: time.Now(),
	}

	result := DB.Conn.Model(&data).Create(&data)

	if result.Error != nil {
		return follows.Domain{}, result.Error
	}

	return follows.Domain{}, nil
}

func (DB *MysqlFollowRepository) Follows(ctx context.Context, domain follows.Domain, my_id int) (follows.Domain, error) {
	var follow Follows
	result := DB.Conn.Model(&follow).Where("user_id = ? AND follower_id = ?", domain.User_id, my_id).
		Updates(Follows{State: true, Followed_at: time.Now()})

	if result.Error != nil {
		return follows.Domain{}, result.Error
	}

	return follow.ToDomain(), nil
}

func (DB *MysqlFollowRepository) Unfollow(ctx context.Context, domain follows.Domain, my_id int) (follows.Domain, error) {
	var follow Follows

	result := DB.Conn.Model(&follow).Where("user_id = ? AND follower_id = ?", domain.User_id, my_id).
		Update("state", false)

	if result.Error != nil {
		return follows.Domain{}, result.Error
	}

	return follow.ToDomain(), nil
}

func (DB *MysqlFollowRepository) GetFollowState(ctx context.Context, domain follows.Domain, my_id int) (follows.Domain, error) {
	var follow Follows

	result := DB.Conn.Where("user_id = ? AND follower_id = ?", domain.User_id, my_id).Find(&follow)

	if result.Error != nil {
		return follows.Domain{}, result.Error
	}

	return follow.ToDomain(), nil
}
