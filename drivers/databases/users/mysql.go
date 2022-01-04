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

func (DB *MysqlUserRepository) GetUser(ctx context.Context) ([]users.Domain, error) {
	var User []Users

	result := DB.Conn.Table("users").Find(&User)

	if result.Error != nil {
		return []users.Domain{}, result.Error
	}
	return ToListDomain(User), nil
}
