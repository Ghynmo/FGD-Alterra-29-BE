package comments

import (
	"context"
	"fgd-alterra-29/business/comments"
	"fmt"

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

func (DB *MysqlCommentRepository) GetCommentProfile(ctx context.Context, id int) ([]comments.Domain, error) {
	var Comment []Comments

	ReplierName := DB.Conn.Table("comments as subcomment").Where("subcomment.id = comments.reply_of").Select("name").
		Joins("join users on subcomment.user_id = users.id")

	result := DB.Conn.Table("comments").Where("comments.user_id = 1").Select("title as Thread, comment, (?) as Name", ReplierName).
		Joins("join threads on comments.thread_id = threads.id").Joins("join users on comments.user_id = users.id").
		Find(&Comment)

	if result.Error != nil {
		return []comments.Domain{}, result.Error
	}

	fmt.Println(ToListDomain(Comment))

	return ToListDomain(Comment), nil
}
