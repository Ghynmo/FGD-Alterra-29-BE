package threads

import (
	"context"
	"fgd-alterra-29/business/threads"

	"gorm.io/gorm"
)

type MysqlThreadRepository struct {
	Conn *gorm.DB
}

func NewMysqlThreadRepository(conn *gorm.DB) threads.Repository {
	return &MysqlThreadRepository{
		Conn: conn,
	}
}

func (DB *MysqlThreadRepository) GetProfileThreads(ctx context.Context, id int) ([]threads.Domain, error) {
	var Thread []Threads

	SQReplier := DB.Conn.Table("comments").Select("name").Joins("join users on comments.user_id = users.id").Where("users.id = comments.user_id").Limit(1)
	Replier := DB.Conn.Table("comments").Select("(?) as name", SQReplier).Where("comments.thread_id = threads.id").Limit(1)

	Comment := DB.Conn.Table("comments").Select("comment").Where("comments.thread_id = threads.id").Limit(1)

	result := DB.Conn.Table("threads").Where("threads.user_id = 1").Select("id, title, (?) as Comment, (?) as RecentReplier", Comment, Replier).
		Order("created_at desc").Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}
