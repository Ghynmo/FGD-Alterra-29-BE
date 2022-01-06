package threadlikes

import (
	"context"
	threadlikes "fgd-alterra-29/business/thread_likes"

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
	var threadLike ThreadLikes

	result := DB.Conn.Model(&threadLike).Create(&domain)

	if result.Error != nil {
		return threadlikes.Domain{}, result.Error
	}

	return threadlikes.Domain{}, nil
}

func (DB *MysqlThreadLikeRepository) Unlike(ctx context.Context, domain threadlikes.Domain) (threadlikes.Domain, error) {
	var threadLike ThreadLikes

	result := DB.Conn.Where("thread_id = ? AND user_id = ?", domain.Thread_id, domain.User_id).Delete(&threadLike)

	if result.Error != nil {
		return threadlikes.Domain{}, result.Error
	}

	return threadlikes.Domain{}, nil
}
