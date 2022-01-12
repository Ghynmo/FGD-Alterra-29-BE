package threadshares

import (
	"context"
	threadshares "fgd-alterra-29/business/thread_shares"
	"time"

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

	data := ThreadShares{
		User_id:   domain.User_id,
		Thread_id: domain.Thread_id,
		Shared_at: time.Now(),
	}

	result := DB.Conn.Model(&data).Create(&domain)

	if result.Error != nil {
		return threadshares.Domain{}, result.Error
	}

	return data.ToDomain(), nil
}
