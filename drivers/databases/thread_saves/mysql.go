package threadsaves

import (
	"context"
	threadsaves "fgd-alterra-29/business/thread_saves"
	"time"

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

func (DB *MysqlThreadSaveRepository) NewSave(ctx context.Context, domain threadsaves.Domain, id int) (threadsaves.Domain, error) {
	data := ThreadSaves{
		User_id:   id,
		Thread_id: domain.Thread_id,
		Saved_at:  time.Now(),
	}
	result := DB.Conn.Model(&data).Create(&data)

	if result.Error != nil {
		return threadsaves.Domain{}, result.Error
	}

	return data.ToDomain(), nil
}

func (DB *MysqlThreadSaveRepository) Save(ctx context.Context, domain threadsaves.Domain, id int) (threadsaves.Domain, error) {
	var TL ThreadSaves
	result := DB.Conn.Model(&TL).Where("thread_id = ? AND user_id = ?", domain.Thread_id, id).
		Updates(ThreadSaves{State: true, Saved_at: time.Now()})

	if result.Error != nil {
		return threadsaves.Domain{}, result.Error
	}

	return TL.ToDomain(), nil
}

func (DB *MysqlThreadSaveRepository) Unsave(ctx context.Context, domain threadsaves.Domain, id int) (threadsaves.Domain, error) {
	var ThreadSave ThreadSaves

	result := DB.Conn.Model(&ThreadSave).Where("thread_id = ? AND user_id = ?", domain.Thread_id, id).
		Update("state", false)

	if result.Error != nil {
		return threadsaves.Domain{}, result.Error
	}

	return ThreadSave.ToDomain(), nil
}

func (DB *MysqlThreadSaveRepository) GetSaveState(ctx context.Context, domain threadsaves.Domain, id int) (threadsaves.Domain, error) {
	var TL ThreadSaves

	result := DB.Conn.Where("thread_id = ? AND user_id = ?", domain.Thread_id, id).Find(&TL)

	if result.Error != nil {
		return threadsaves.Domain{}, result.Error
	}

	return TL.ToDomain(), nil
}
