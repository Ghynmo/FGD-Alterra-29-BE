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

func (DB *MysqlBadgeRepository) GetBadgesByPoint(ctx context.Context, point int) ([]badges.Domain, error) {
	var Badge []Badges

	result := DB.Conn.Table("badges").Where("requirement_point < ?", point).Find(&Badge)

	if result.Error != nil {
		return []badges.Domain{}, result.Error
	}
	return ToListDomain(Badge), nil
}
