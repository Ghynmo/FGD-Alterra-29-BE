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
