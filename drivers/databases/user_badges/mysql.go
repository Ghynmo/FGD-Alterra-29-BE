package userbadges

import (
	"context"
	userbadges "fgd-alterra-29/business/user_badges"

	"gorm.io/gorm"
)

type MysqlUserBadgeRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserBadgeRepository(conn *gorm.DB) userbadges.Repository {
	return &MysqlUserBadgeRepository{
		Conn: conn,
	}
}

func (DB *MysqlUserBadgeRepository) GetUserBadge(ctx context.Context, id int) ([]userbadges.Domain, error) {
	var UserBadges []UserBadges

	result := DB.Conn.Table("user_badges").Select("badge").Where("user_badges.user_id = ?", id).
		Joins("join badges on user_badges.badge_id = badges.id").
		Find(&UserBadges)

	if result.Error != nil {
		return []userbadges.Domain{}, result.Error
	}

	return ToListDomain(UserBadges), nil
}

func (DB *MysqlUserBadgeRepository) CheckGetBadge(ctx context.Context, user_id int, badge_id int) (userbadges.Domain, error) {
	var UserBadges UserBadges

	result := DB.Conn.Table("user_badges").Where("user_id = ? AND badge_id = ?", user_id, badge_id).
		Find(&UserBadges)

	if result.Error != nil {
		return userbadges.Domain{}, result.Error
	}

	return UserBadges.ToDomain(), nil
}

func (DB *MysqlUserBadgeRepository) CreatenewRecord(ctx context.Context, user_id int, badge_id int) (userbadges.Domain, error) {

	data := UserBadges{
		User_id:  user_id,
		Badge_id: badge_id,
	}

	result := DB.Conn.Model(&data).Create(&data)

	if result.Error != nil {
		return userbadges.Domain{}, result.Error
	}

	return data.ToDomain(), nil
}
