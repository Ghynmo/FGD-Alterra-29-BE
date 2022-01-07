package comments

import (
	"context"
	"fgd-alterra-29/business/comments"

	"gorm.io/gorm"
)

type MysqlCommentRepository struct {
	Conn *gorm.DB
}

func NewMysqlCommentRepository(conn *gorm.DB) comments.Repository {
	return &MysqlCommentRepository{
		Conn: conn,
	}
}

func (DB *MysqlCommentRepository) GetPostsByComment(ctx context.Context, comment string) ([]comments.Domain, error) {
	var Comment []Comments
	var NewComment = ("%" + comment + "%")

	result := DB.Conn.Table("comments").Select("comments.id, name, photo_url as Photo, comment, comments.created_at").
		Where("comment LIKE ?", NewComment).Joins("join users on comments.user_id = users.id").
		Find(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}

	return ToListDomain(Comment), nil
}

func (DB *MysqlCommentRepository) GetCommentProfile(ctx context.Context, id int) ([]comments.Domain, error) {
	var Comment []Comments

	result := DB.Conn.Table("comments").Where("comments.user_id = 1").Select("title as Thread, comment, name").
		Joins("join threads on comments.thread_id = threads.id").Joins("join users on comments.user_id = users.id").
		Find(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}

	return ToListDomain(Comment), nil
}

func (DB *MysqlCommentRepository) GetPostQuantity(ctx context.Context) (comments.Domain, error) {
	var Comment Comments

	result := DB.Conn.Table("comments").Select("count(id) as Q_Post").Find(&Comment)

	if result.Error != nil {
		return comments.Domain{}, result.Error
	}

	return Comment.ToDomain(), nil
}

func (DB *MysqlCommentRepository) GetPosts(ctx context.Context) ([]comments.Domain, error) {
	var Comment []Comments

	result := DB.Conn.Table("comments").Select("comments.id, name, photo_url as Photo, comment, comments.created_at").
		Joins("join users on comments.user_id = users.id").Order("comments.created_at desc").
		Find(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}

	return ToListDomain(Comment), nil
}

func (DB *MysqlCommentRepository) DeletePost(ctx context.Context, id int) (comments.Domain, error) {
	var Comment Comments

	result := DB.Conn.Delete(&Comment, id)

	if result.Error != nil {
		return comments.Domain{}, result.Error
	}

	return Comment.ToDomain(), nil
}
