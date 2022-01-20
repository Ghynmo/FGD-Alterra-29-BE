package users

import (
	"context"
	"fgd-alterra-29/business/users"
	"time"

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

	data := Users{
		Name:       domain.Name,
		Email:      domain.Email,
		Password:   domain.Password,
		Created_at: time.Now(),
	}

	result := DB.Conn.Model(&data).Create(&data)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return data.ToDomain(), nil
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
	Q_Follower := DB.Conn.Table("follows").Select("count(follower_id)").Where("user_id = ? AND users.status = ?", id, "active").Joins("join users on follows.follower_id = users.id").Group("user_id")
	Q_Following := DB.Conn.Table("follows").Select("count(user_id)").Where("follower_id = ? AND users.status = ?", id, "active").Joins("join users on follows.user_id = users.id").Group("follower_id")
	Q_Post := DB.Conn.Table("comments").Select("count(comment)").Where("user_id = ? AND comments.active = 1", id).Group("user_id")
	Q_Thread := DB.Conn.Table("threads").Select("count(title)").Where("user_id = ? AND threads.active = 1", id).Group("user_id")

	result := DB.Conn.Table("users").Where("users.id = ? AND users.status = ?", id, "active").
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

	result := DB.Conn.Table("users").Select("id, name, photo_url, email, phone, bio, address").
		Where("id = (?) AND users.status = ?", id, "active").Find(&User)

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

	result2nd, err2 := DB.GetProfileSetting(ctx, id)

	if err2 != nil {
		return users.Domain{}, result.Error
	}

	return result2nd, nil
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

func (DB *MysqlUserRepository) GetBannedState(ctx context.Context, id int) (users.Domain, error) {
	var User Users

	result := DB.Conn.Select("status").Where("user_id = ?", id).Find(&User)

	if result.Error != nil {
		return users.Domain{}, result.Error
	}

	return User.ToDomain(), nil
}

func (DB *MysqlUserRepository) CheckUsername(ctx context.Context, username string) (bool, error) {
	var Uname string

	result := DB.Conn.Table("users").Where("name = ?", username).Select("name").Row()
	result.Scan(&Uname)

	if result.Err() != nil {
		return false, result.Err()
	}

	if Uname == username {
		return true, nil
	}

	return false, nil
}

func (DB *MysqlUserRepository) CheckEmail(ctx context.Context, email string) (bool, error) {
	var Email string

	result := DB.Conn.Table("users").Where("email = ?", email).Select("email").Row()
	result.Scan(&Email)

	if result.Err() != nil {
		return false, result.Err()
	}

	if Email == email {
		return true, nil
	}

	return false, nil
}
