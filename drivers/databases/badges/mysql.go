package badges

import (
	"context"
	"fgd-alterra-29/business/badges"

	"gorm.io/gorm"
)

type MysqlBadgeRepository struct {
	Conn *gorm.DB
}

func NewMysqlBadgeRepository(conn *gorm.DB) badges.Repository {
	return &MysqlBadgeRepository{
		Conn: conn,
	}
}

func (DB *MysqlBadgeRepository) GetBadgesIdByThread(ctx context.Context, mythread_qty int) (int, error) {

	var NewBadge_id int

	row := DB.Conn.Table("badges").Select("id").Where("requirement_thread <= ?", mythread_qty).Order("requirement_thread desc").Limit(1).Row()
	row.Scan(&NewBadge_id)

	if row.Err() != nil {
		return 0, row.Err()
	}
	return NewBadge_id, nil
}

func (DB *MysqlBadgeRepository) GetBadgesByUser(ctx context.Context, id int) ([]badges.Domain, error) {
	var Badges []Badges

	result := DB.Conn.Table("badges").Select("badge, badge_url").Where("user_id = ?", id).
		Joins("join user_badges on badges.id = user_badges.badge_id").Find(&Badges)

	if result.Error != nil {
		return []badges.Domain{}, result.Error
	}
	return ToListDomain(Badges), nil
}

func (DB *MysqlBadgeRepository) CreateBadge(ctx context.Context, domain badges.Domain) (badges.Domain, error) {
	var Badge Badges

	result := DB.Conn.Model(&Badge).Create(&domain)

	if result.Error != nil {
		return badges.Domain{}, result.Error
	}
	return Badge.ToDomain(), nil
}

func (DB *MysqlBadgeRepository) ActivateBadge(ctx context.Context, domain badges.Domain) (badges.Domain, error) {
	var Badge Badges

	result := DB.Conn.Table("badges").Where("badges.requirement_thread < ?").
		Find(&Badge)

	if result.Error != nil {
		return badges.Domain{}, result.Error
	}
	return Badge.ToDomain(), nil
}

func (DB *MysqlBadgeRepository) UnactivateBadge(ctx context.Context, domain badges.Domain) (badges.Domain, error) {
	var Badge Badges

	result := DB.Conn.Table("badges").Where("badges.requirement_thread < ?").
		Find(&Badge)

	if result.Error != nil {
		return badges.Domain{}, result.Error
	}
	return Badge.ToDomain(), nil
}
