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

func (DB *MysqlBadgeRepository) GetBadgesByUser(ctx context.Context, id int) ([]badges.Domain, error) {
	var Badge []Badges
	var points int

	row := DB.Conn.Table("users").Select("points").Where("users.id = ?", id).Row()
	row.Scan(&points)

	result := DB.Conn.Table("badges").Where("badges.requirement_point < ?", points).
		Find(&Badge)

	if result.Error != nil {
		return []badges.Domain{}, result.Error
	}
	return ToListDomain(Badge), nil
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

	result := DB.Conn.Table("badges").Where("badges.requirement_point < ?").
		Find(&Badge)

	if result.Error != nil {
		return badges.Domain{}, result.Error
	}
	return Badge.ToDomain(), nil
}

func (DB *MysqlBadgeRepository) UnactivateBadge(ctx context.Context, domain badges.Domain) (badges.Domain, error) {
	var Badge Badges

	result := DB.Conn.Table("badges").Where("badges.requirement_point < ?").
		Find(&Badge)

	if result.Error != nil {
		return badges.Domain{}, result.Error
	}
	return Badge.ToDomain(), nil
}
