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

	result := DB.Conn.Table("comments").Select("comments.id, name, photo_url as Photo, comment, active, comments.created_at").
		Where("comment LIKE ?", NewComment).Joins("join users on comments.user_id = users.id").
		Find(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}

	return ToListDomain(Comment), nil
}

func (DB *MysqlCommentRepository) GetCommentReply(ctx context.Context, id int) ([]comments.Domain, error) {
	var Comment []Comments

	result := DB.Conn.Table("comments").Where("reply_of = ?", id).Select("comments.id, name, photo_url, comment").
		Joins("join users on comments.reply_of = users.id").
		Find(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}

	return ToListDomain(Comment), nil
}

func (DB *MysqlCommentRepository) GetCommentProfile(ctx context.Context, id int) ([]comments.Domain, error) {
	var Comment []Comments

	result := DB.Conn.Table("comments").Where("comments.user_id = 1 AND comments.active = 1").Select("comments.id, title as Thread, comment, name").
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

	result := DB.Conn.Table("comments").Select("comments.id, name, photo_url as Photo, comment, active, comments.created_at").
		Joins("join users on comments.user_id = users.id").Order("comments.created_at desc").
		Find(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}

	return ToListDomain(Comment), nil
}

func (DB *MysqlCommentRepository) GetCommentByThread(ctx context.Context, id int) ([]comments.Domain, error) {
	var Comment []Comments

	result := DB.Conn.Table("comments").Where("thread_id = 3").Find(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}

	return ToListDomain(Comment), nil
}

func (DB *MysqlCommentRepository) UnactivatingPost(ctx context.Context, id int) (comments.Domain, error) {
	var Comment Comments

	result := DB.Conn.Model(&Comment).Where("comments.id = ?", id).Update("active", "false")

	if result.Error != nil {
		return comments.Domain{}, result.Error
	}

	return Comment.ToDomain(), nil
}

func (DB *MysqlCommentRepository) ActivatingPost(ctx context.Context, id int) (comments.Domain, error) {
	var Comment Comments

	result := DB.Conn.Model(&Comment).Where("comments.id = ?", id).Update("active", "true")

	if result.Error != nil {
		return comments.Domain{}, result.Error
	}

	return Comment.ToDomain(), nil
}

func (DB *MysqlCommentRepository) CreateComment(ctx context.Context, domain comments.Domain) (comments.Domain, error) {
	var Comment Comments

	if domain.ReplyOf == 0 {
		result := DB.Conn.Exec("INSERT INTO comments (thread_id, user_id, comment, reply_of) VALUES (?, ?, ?, NULL)",
			domain.Thread_id, domain.User_id, domain.Comment)
		if result.Error != nil {
			return comments.Domain{}, result.Error
		}
	} else {
		result := DB.Conn.Exec("INSERT INTO comments (thread_id, user_id, comment, reply_of) VALUES (?, ?, ?, ?)",
			domain.Thread_id, domain.User_id, domain.Comment, domain.ReplyOf)
		if result.Error != nil {
			return comments.Domain{}, result.Error
		}
	}

	return Comment.ToDomain(), nil
}
