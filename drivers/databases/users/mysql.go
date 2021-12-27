package users

import (
	"context"
	"fgd-alterra-29/business/users"

	"gorm.io/gorm"
)

type MysqlUserRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserRepository(conn *gorm.DB) users.Repository {
	return &MysqlUserRepository{
		Conn: conn,
	}
}

func (DB *MysqlUserRepository) GetUsers(ctx context.Context) ([]users.Domain, error) {
	var User []Users

	result := DB.Conn.Table("users").Select("status, photo_url, name, email").Find(&User)

	if result.Error != nil {
		return []users.Domain{}, result.Error
	}
	return ToListDomain(User), nil
}

func (DB *MysqlUserRepository) GetProfile(ctx context.Context, id int) (users.Domain, error) {
	var User Users
	Reputation := DB.Conn.Table("reputations").Select("reputation").Where("users.id = 1").Joins("join users on reputations.id = users.reputation_id")
	Q_Follower := DB.Conn.Table("follows").Select("count(follower_id)").Where("user_id = 1").Group("user_id")
	Q_Following := DB.Conn.Table("follows").Select("count(user_id)").Where("follower_id = 1").Group("follower_id")
	Q_Post := DB.Conn.Table("comments").Select("count(comment)").Where("user_id = 1").Group("user_id")
	Q_Thread := DB.Conn.Table("threads").Select("count(title)").Where("user_id = 1").Group("user_id")

	result := DB.Conn.Table("users").
		Select("*, (?) as Q_Followers, (?) as Q_Following, (?) as Q_Post, (?) as Q_Thread, (?) as Reputation",
			Q_Follower, Q_Following, Q_Post, Q_Thread, Reputation).
		Find(&User)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return User.ToDomain(), nil
}

func (DB *MysqlUserRepository) GetUsersQuantity(ctx context.Context) (users.Domain, error) {
	var User Users

	result := DB.Conn.Table("users").Select("count(id) as Q_User").Find(&User)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return User.ToDomain(), nil
}

func (DB *MysqlUserRepository) GetUserSetting(ctx context.Context, id int) (users.Domain, error) {
	var User Users

	result := DB.Conn.Table("users").Select("id, name, photo_url, email, phone").Where("id = (?)", id).Find(&User)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return User.ToDomain(), nil
}
