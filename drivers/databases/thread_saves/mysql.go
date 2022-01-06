package threadsaves

import (
	"context"
	threadsaves "fgd-alterra-29/business/thread_saves"

	"gorm.io/gorm"
)

type MysqlThreadSaveRepository struct {
	Conn *gorm.DB
}

func NewMysqlThreadSaveRepository(conn *gorm.DB) threadsaves.Repository {
	return &MysqlThreadSaveRepository{
		Conn: conn,
	}
}

func (DB *MysqlThreadSaveRepository) SaveThread(ctx context.Context, domain threadsaves.Domain) (threadsaves.Domain, error) {
	var threadSave ThreadSaves

	result := DB.Conn.Model(&threadSave).Create(&domain)

	if result.Error != nil {
		return threadsaves.Domain{}, result.Error
	}

	return threadsaves.Domain{}, nil
}

func (DB *MysqlThreadSaveRepository) UnsaveThread(ctx context.Context, domain threadsaves.Domain) (threadsaves.Domain, error) {
	var threadSave ThreadSaves

	result := DB.Conn.Where("thread_id = ? AND user_id = ?", domain.Thread_id, domain.User_id).Delete(&threadSave)

	if result.Error != nil {
		return threadsaves.Domain{}, result.Error
	}

	return threadsaves.Domain{}, nil
}
