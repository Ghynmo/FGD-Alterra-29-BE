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

func (DB *MysqlCommentLikeRepository) Like(ctx context.Context, domain commentlikes.Domain) (commentlikes.Domain, error) {

	data := CommentLikes{
		Comment_id: domain.Comment_id,
		Liker_id:   domain.Liker_id,
		Liked_at:   time.Now(),
	}

	result := DB.Conn.Model(&data).Create(&domain)

	if result.Error != nil {
		return commentlikes.Domain{}, result.Error
	}

	return data.ToDomain(), nil
}

func (DB *MysqlCommentLikeRepository) Unlike(ctx context.Context, domain commentlikes.Domain) (commentlikes.Domain, error) {
	var CommentLike CommentLikes

	result := DB.Conn.Where("comment_id = ? AND liker_id = ?", domain.Comment_id, domain.Liker_id).Delete(&CommentLike)

	if result.Error != nil {
		return commentlikes.Domain{}, result.Error
	}

	return CommentLike.ToDomain(), nil
}
