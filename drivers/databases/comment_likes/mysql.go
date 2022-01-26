package commentlikes

import (
	"context"
	commentlikes "fgd-alterra-29/business/comment_likes"
	"time"

	"gorm.io/gorm"
)

type MysqlCommentLikeRepository struct {
	Conn *gorm.DB
}

func NewMysqlCommentLikeRepository(conn *gorm.DB) commentlikes.Repository {
	return &MysqlCommentLikeRepository{
		Conn: conn,
	}
}

func (DB *MysqlCommentLikeRepository) NewLike(ctx context.Context, domain commentlikes.Domain, id int) (commentlikes.Domain, int, error) {
	var thread_user_id int

	data := CommentLikes{
		Comment_id: domain.Comment_id,
		Liker_id:   id,
		Liked_at:   time.Now(),
	}
	row := DB.Conn.Table("comments").Where("id = ?", domain.Comment_id).Select("user_id").Row()
	row.Scan(&thread_user_id)
	result := DB.Conn.Model(&data).Create(&data)

	if result.Error != nil {
		return commentlikes.Domain{}, 0, result.Error
	}

	return data.ToDomain(), thread_user_id, nil
}

func (DB *MysqlCommentLikeRepository) Like(ctx context.Context, domain commentlikes.Domain, id int) (commentlikes.Domain, int, error) {
	var CL CommentLikes
	var thread_user_id int
	result := DB.Conn.Model(&CL).Where("comment_id = ? AND liker_id = ?", domain.Comment_id, id).
		Updates(CommentLikes{State: true, Liked_at: time.Now()})
	row := DB.Conn.Table("comments").Where("id = ?", domain.Comment_id).Select("user_id").Row()
	row.Scan(&thread_user_id)

	if result.Error != nil {
		return commentlikes.Domain{}, 0, result.Error
	}
	return CL.ToDomain(), thread_user_id, nil
}

func (DB *MysqlCommentLikeRepository) Unlike(ctx context.Context, domain commentlikes.Domain, id int) (commentlikes.Domain, int, error) {
	var CommentLike CommentLikes
	var thread_user_id int

	result := DB.Conn.Model(&CommentLike).Where("comment_id = ? AND liker_id = ?", domain.Comment_id, id).
		Update("state", false)
	row := DB.Conn.Table("comments").Where("id = ?", domain.Comment_id).Select("user_id").Row()
	row.Scan(&thread_user_id)
	if result.Error != nil {
		return commentlikes.Domain{}, 0, result.Error
	}

	return CommentLike.ToDomain(), thread_user_id, nil
}

func (DB *MysqlCommentLikeRepository) GetLikeState(ctx context.Context, domain commentlikes.Domain, id int) (commentlikes.Domain, error) {
	var CL CommentLikes

	result := DB.Conn.Where("comment_id = ? AND liker_id = ?", domain.Comment_id, id).Find(&CL)

	if result.Error != nil {
		return commentlikes.Domain{}, result.Error
	}

	return CL.ToDomain(), nil
}
