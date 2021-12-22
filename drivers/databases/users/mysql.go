package users

import (
	"context"
	b_threads "fgd-alterra-29/business/threads"
	"fgd-alterra-29/business/users"
	threads "fgd-alterra-29/drivers/databases/threads"

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

	result := DB.Conn.Table("users").Find(&User)

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

func (DB *MysqlUserRepository) GetCategoryActive(ctx context.Context, id int) ([]b_threads.Domain, error) {
	var Threads []threads.Threads

	result := DB.Conn.Table("threads").Select("count(title) as Q_Title, category").Where("threads.user_id = 1").
		Joins("join categories on threads.category_id = categories.id").Group("category_id").
		Find(&Threads)

	if result.Error != nil {
		return []b_threads.Domain{}, result.Error
	}

	return threads.ToListDomain(Threads), nil
}
