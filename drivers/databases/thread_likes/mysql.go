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

func (DB *MysqlThreadLikeRepository) NewLike(ctx context.Context, domain threadlikes.Domain, id int) (threadlikes.Domain, int, error) {
	var thread_user_id int
	data := ThreadLikes{
		User_id:   id,
		Thread_id: domain.Thread_id,
		Liked_at:  time.Now(),
	}
	result := DB.Conn.Model(&data).Create(&data)
	row := DB.Conn.Table("threads").Where("id = ?", domain.Thread_id).Select("user_id").Row()
	row.Scan(&thread_user_id)

	if result.Error != nil {
		return threadlikes.Domain{}, 0, result.Error
	}

	return data.ToDomain(), thread_user_id, nil
}

func (DB *MysqlThreadLikeRepository) Like(ctx context.Context, domain threadlikes.Domain, id int) (threadlikes.Domain, int, error) {
	var TL ThreadLikes
	var thread_user_id int
	result := DB.Conn.Model(&TL).Where("thread_id = ? AND user_id = ?", domain.Thread_id, id).
		Updates(ThreadLikes{State: true, Liked_at: time.Now()})
	row := DB.Conn.Table("threads").Where("id = ?", domain.Thread_id).Select("user_id").Row()
	row.Scan(&thread_user_id)

	if result.Error != nil {
		return threadlikes.Domain{}, 0, result.Error
	}

	return TL.ToDomain(), thread_user_id, nil
}

func (DB *MysqlThreadLikeRepository) Unlike(ctx context.Context, domain threadlikes.Domain, id int) (threadlikes.Domain, int, error) {
	var ThreadLike ThreadLikes
	var thread_user_id int

	result := DB.Conn.Model(&ThreadLike).Where("thread_id = ? AND user_id = ?", domain.Thread_id, id).
		Update("state", false)
	row := DB.Conn.Table("threads").Where("id = ?", domain.Thread_id).Select("user_id").Row()
	row.Scan(&thread_user_id)

	if result.Error != nil {
		return threadlikes.Domain{}, 0, result.Error
	}

	return ThreadLike.ToDomain(), thread_user_id, nil
}

func (DB *MysqlThreadLikeRepository) GetLikeState(ctx context.Context, domain threadlikes.Domain, id int) (threadlikes.Domain, error) {
	var TL ThreadLikes

	result := DB.Conn.Where("thread_id = ? AND user_id = ?", domain.Thread_id, id).Find(&TL)

	if result.Error != nil {
		return threadlikes.Domain{}, result.Error
	}

	return TL.ToDomain(), nil
}
