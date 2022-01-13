package threadlikes

import (
	"context"
	threadlikes "fgd-alterra-29/business/thread_likes"
	"time"

	"gorm.io/gorm"
)

type MysqlThreadLikeRepository struct {
	Conn *gorm.DB
}

func NewMysqlThreadLikeRepository(conn *gorm.DB) threadlikes.Repository {
	return &MysqlThreadLikeRepository{
		Conn: conn,
	}
}

func (DB *MysqlThreadLikeRepository) NewLike(ctx context.Context, domain threadlikes.Domain) (threadlikes.Domain, error) {
	data := ThreadLikes{
		User_id:   domain.User_id,
		Thread_id: domain.Thread_id,
		Liked_at:  time.Now(),
	}
	result := DB.Conn.Model(&data).Create(&data)

	if result.Error != nil {
		return threadlikes.Domain{}, result.Error
	}

	return data.ToDomain(), nil
}

func (DB *MysqlThreadLikeRepository) Like(ctx context.Context, domain threadlikes.Domain) (threadlikes.Domain, error) {
	var TL ThreadLikes
	result := DB.Conn.Model(&TL).Where("thread_id = ? AND user_id = ?", domain.Thread_id, domain.User_id).
		Updates(ThreadLikes{State: true, Liked_at: time.Now()})

	if result.Error != nil {
		return threadlikes.Domain{}, result.Error
	}

	return TL.ToDomain(), nil
}

func (DB *MysqlThreadLikeRepository) Unlike(ctx context.Context, domain threadlikes.Domain) (threadlikes.Domain, error) {
	var ThreadLike ThreadLikes

	result := DB.Conn.Model(&ThreadLike).Where("thread_id = ? AND user_id = ?", domain.Thread_id, domain.User_id).
		Update("state", false)

	if result.Error != nil {
		return threadlikes.Domain{}, result.Error
	}

	return ThreadLike.ToDomain(), nil
}

func (DB *MysqlThreadLikeRepository) GetLikeState(ctx context.Context, domain threadlikes.Domain) (threadlikes.Domain, error) {
	var TL ThreadLikes

	result := DB.Conn.Where("thread_id = ? AND user_id = ?", domain.Thread_id, domain.User_id).Find(&TL)

	if result.Error != nil {
		return threadlikes.Domain{}, result.Error
	}

	return TL.ToDomain(), nil
}
