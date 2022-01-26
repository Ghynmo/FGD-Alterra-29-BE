package reputations

import (
	"context"
	"fgd-alterra-29/business/reputations"

	"gorm.io/gorm"
)

type MysqlReputationRepository struct {
	Conn *gorm.DB
}

func NewMysqlReputationRepository(conn *gorm.DB) reputations.Repository {
	return &MysqlReputationRepository{
		Conn: conn,
	}
}

func (DB *MysqlReputationRepository) CreateReputation(ctx context.Context, domain reputations.Domain) (reputations.Domain, error) {
	var reputation Reputations
	result := DB.Conn.Model(&reputation).Create(&domain)

	if result.Error != nil {
		return reputations.Domain{}, result.Error
	}

	return reputation.ToDomain(), nil
}

func (DB *MysqlReputationRepository) GetReputationByUser(ctx context.Context, id int) (reputations.Domain, error) {
	var Badge Reputations
	var points int

	row := DB.Conn.Table("users").Select("points").Where("users.id = ?", id).Row()
	row.Scan(&points)

	result := DB.Conn.Table("reputations").Where("like_points <= ?", points).Order("like_points desc").
		Find(&Badge)

	if result.Error != nil {
		return reputations.Domain{}, result.Error
	}
	return Badge.ToDomain(), nil
}
