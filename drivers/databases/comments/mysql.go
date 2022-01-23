package comments

import (
	"context"
	"fgd-alterra-29/business/comments"
	"time"

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
		Where("comment LIKE ? AND comments.active = 1", NewComment).Joins("join users on comments.user_id = users.id").
		Order("comments.created_at desc").Find(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}

	return ToListDomain(Comment), nil
}

func (DB *MysqlCommentRepository) GetCommentByThread(ctx context.Context, thread_id int, my_id int) ([]comments.Domain, error) {
	var Comment []Comments

	result := DB.Conn.Raw("SELECT comments.id, name, photo_url, comment, state AS LikeState FROM comments LEFT JOIN (SELECT * FROM comment_likes WHERE comment_likes.liker_id = ?) AS B ON comments.id = B.comment_id JOIN users ON comments.user_id = users.id WHERE (thread_id = ? AND comments.active = ?)",
		my_id, thread_id, "true").Scan(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}
	return ToListDomain(Comment), nil
}

func (DB *MysqlCommentRepository) GetCommentReply(ctx context.Context, id int, reply_of int) ([]comments.Domain, error) {
	var Comment []Comments

	result := DB.Conn.Raw("SELECT comments.id, name, photo_url, comment, state AS LikeState FROM comments LEFT JOIN (SELECT * FROM comment_likes WHERE comment_likes.liker_id = ?) AS B ON comments.id = B.comment_id JOIN users ON comments.user_id = users.id WHERE (reply_of = ? AND comments.active = ?)",
		id, reply_of, "true").Scan(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}
	return ToListDomain(Comment), nil
}

func (DB *MysqlCommentRepository) GetCommentProfile(ctx context.Context, id int) ([]comments.Domain, error) {
	var Comment []Comments

	ReplierName := DB.Conn.Table("comments as subcomment").Where("subcomment.id = comments.reply_of").Select("name").
		Joins("join users on subcomment.user_id = users.id")

	result := DB.Conn.Table("comments").Where("comments.user_id = ? AND comments.active = 1", id).Select("title as Thread, comment, (?) as Name", ReplierName).
		Joins("join threads on comments.thread_id = threads.id").Joins("join users on comments.user_id = users.id").
		Order("comments.created_at desc").Find(&Comment)

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

func (DB *MysqlCommentRepository) UnactivatingPost(ctx context.Context, id int) (comments.Domain, error) {
	var Comment Comments

	result := DB.Conn.Model(&Comment).Where("comments.id = ?", id).Update("active", false)

	if result.Error != nil {
		return comments.Domain{}, result.Error
	}

	return Comment.ToDomain(), nil
}

func (DB *MysqlCommentRepository) ActivatingPost(ctx context.Context, id int) (comments.Domain, error) {
	var Comment Comments

	result := DB.Conn.Model(&Comment).Where("comments.id = ?", id).Update("active", true)

	if result.Error != nil {
		return comments.Domain{}, result.Error
	}

	return Comment.ToDomain(), nil
}

func (DB *MysqlCommentRepository) CreateComment(ctx context.Context, domain comments.Domain, id int) (comments.Domain, error) {
	var Comment Comments

	if domain.ReplyOf == 0 {
		result := DB.Conn.Exec("INSERT INTO comments (thread_id, user_id, comment, reply_of, created_at) VALUES (?, ?, ?, NULL, ?)",
			domain.Thread_id, id, domain.Comment, time.Now())
		if result.Error != nil {
			return comments.Domain{}, result.Error
		}
	} else {
		result := DB.Conn.Exec("INSERT INTO comments (thread_id, user_id, comment, reply_of, created_at) VALUES (?, ?, ?, ?, ?)",
			domain.Thread_id, id, domain.Comment, domain.ReplyOf, time.Now())
		if result.Error != nil {
			return comments.Domain{}, result.Error
		}
	}

	return Comment.ToDomain(), nil
}

// func (DB *MysqlCommentRepository) GetCommentParent(ctx context.Context, id int) ([]comments.Domain, error) {
// 	var Comment []Comments
// 	var parent_id int

// 	row := DB.Conn.Table("comments").Select("reply_of").Where("comments.id = ?", id).Row()
// 	row.Scan(&parent_id)

// 	result := DB.Conn.Table("comments").Where("comments.id = ? AND comments.active = 1", id).Select("comments.id, name, photo_url, comment").
// 		Joins("join users on comments.user_id = users.id").Order("comments.created_at desc").
// 		Find(&Comment)

// 	if result.Error != nil {
// 		return []comments.Domain{}, result.Error
// 	}
// 	return ToListDomain(Comment), nil
// }
