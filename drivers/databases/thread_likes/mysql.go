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

func (DB *MysqlThreadLikeRepository) Like(ctx context.Context, domain threadlikes.Domain) (threadlikes.Domain, error) {
	data := ThreadLikes{
		User_id:   domain.User_id,
		Thread_id: domain.Thread_id,
		Liked_at:  time.Now(),
	}
	result := DB.Conn.Model(&data).Create(&domain)

	if result.Error != nil {
		return threadlikes.Domain{}, result.Error
	}

	return data.ToDomain(), nil
}

func (DB *MysqlThreadLikeRepository) Unlike(ctx context.Context, domain threadlikes.Domain) (threadlikes.Domain, error) {
	var threadLike ThreadLikes

	result := DB.Conn.Where("thread_id = ? AND user_id = ?", domain.Thread_id, domain.User_id).Delete(&threadLike)

	if result.Error != nil {
		return threadlikes.Domain{}, result.Error
	}

	return threadlikes.Domain{}, nil
}
