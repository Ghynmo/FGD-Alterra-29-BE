package threadshares

import (
	"context"
	threadshares "fgd-alterra-29/business/thread_shares"

	"gorm.io/gorm"
)

type MysqlThreadShareRepository struct {
	Conn *gorm.DB
}

func NewMysqlThreadShareRepository(conn *gorm.DB) threadshares.Repository {
	return &MysqlThreadShareRepository{
		Conn: conn,
	}
}

func (DB *MysqlThreadShareRepository) ThreadShare(ctx context.Context, domain threadshares.Domain) (threadshares.Domain, error) {
	var threadShare ThreadShares

	result := DB.Conn.Model(&threadShare).Create(&domain)

	if result.Error != nil {
		return threadshares.Domain{}, result.Error
	}

	return threadshares.Domain{}, nil
}
