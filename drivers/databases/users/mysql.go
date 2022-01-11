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

func (DB *MysqlUserRepository) Register(ctx context.Context, domain users.Domain) (users.Domain, error) {
	User := FromDomain(domain)
	result := DB.Conn.Create(&User)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return User.ToDomain(), nil
}

func (DB *MysqlUserRepository) Login(ctx context.Context, domain users.Domain) (users.Domain, error) {
	User := FromDomain(domain)

	result := DB.Conn.First(&User, "email = (?)", User.Email)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return User.ToDomain(), nil
}

func (DB *MysqlUserRepository) GetUsers(ctx context.Context) ([]users.Domain, error) {
	var User []Users

	result := DB.Conn.Table("users").Select("id, status, photo_url, name, email").Find(&User)

	if result.Error != nil {
		return []users.Domain{}, result.Error
	}
	return ToListDomain(User), nil
}

func (DB *MysqlUserRepository) GetUsersByName(ctx context.Context, name string) ([]users.Domain, error) {
	var User []Users
	var NewName = ("%" + name + "%")

	result := DB.Conn.Table("users").Select("id, status, photo_url, name, email").Where("name LIKE ?", NewName).Find(&User)

	if result.Error != nil {
		return []users.Domain{}, result.Error
	}
	return ToListDomain(User), nil
}

func (DB *MysqlUserRepository) GetProfile(ctx context.Context, id int) (users.Domain, error) {
	var User Users
	Reputation := DB.Conn.Table("reputations").Select("reputation").Where("users.id = ?", id).Joins("join users on reputations.id = users.reputation_id")
	Q_Follower := DB.Conn.Table("follows").Select("count(follower_id)").Where("user_id = ?", id).Group("user_id")
	Q_Following := DB.Conn.Table("follows").Select("count(user_id)").Where("follower_id = ?", id).Group("follower_id")
	Q_Post := DB.Conn.Table("comments").Select("count(comment)").Where("user_id = ?", id).Group("user_id")
	Q_Thread := DB.Conn.Table("threads").Select("count(title)").Where("user_id = ?", id).Group("user_id")

	result := DB.Conn.Table("users").Where("users.id = ?", id).
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

func (DB *MysqlUserRepository) GetProfileSetting(ctx context.Context, id int) (users.Domain, error) {
	var User Users

	result := DB.Conn.Table("users").Select("id, name, photo_url, email, phone, bio, address").Where("id = (?)", id).Find(&User)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return User.ToDomain(), nil
}

func (DB *MysqlUserRepository) BannedUser(ctx context.Context, id int) (users.Domain, error) {
	var User Users

	result := DB.Conn.Model(&User).Where("users.id = (?)", id).Update("status", "banned")

	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return User.ToDomain(), nil
}

func (DB *MysqlUserRepository) UnbannedUser(ctx context.Context, id int) (users.Domain, error) {
	var User Users

	result := DB.Conn.Model(&User).Where("users.id = (?)", id).Update("status", "active")

	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return User.ToDomain(), nil
}

func (DB *MysqlUserRepository) UpdateProfile(ctx context.Context, domain users.Domain, id int) (users.Domain, error) {
	var User Users

	result := DB.Conn.Model(&User).Where("id = (?)", id).Updates(domain)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}
	return User.ToDomain(), nil
}
