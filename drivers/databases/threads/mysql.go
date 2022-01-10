package threads

import (
	"context"
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/drivers/databases/comments"
	"fmt"

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

func (DB *MysqlThreadRepository) GetThreadsByTitle(ctx context.Context, title string) ([]threads.Domain, error) {
	var Thread []Threads
	var NewTitle = ("%" + title + "%")

	result := DB.Conn.Table("threads").Select("threads.id, name, photo_url as Photo, title, category, threads.created_at").
		Joins("join users on threads.user_id = users.id").Joins("join categories on threads.category_id = categories.id").
		Where("title LIKE ?", NewTitle).Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetProfileThreads(ctx context.Context, id int) ([]threads.Domain, error) {
	var Thread []Threads

	Comment := DB.Conn.Table("comments").Select("comment").Where("comments.thread_id = threads.id").Order("created_at desc").Limit(1)
	Replier := DB.Conn.Table("comments").Select("name").Joins("join users on comments.user_id = users.id").
		Where("comments.thread_id = threads.id").Order("comments.created_at desc").Limit(1)

	result := DB.Conn.Table("threads").Where("threads.user_id = ?", id).Select("id, title, (?) as Comment, (?) as RecentReplier", Comment, Replier).
		Order("created_at desc").Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetThreadQuantity(ctx context.Context) (threads.Domain, error) {
	var Thread Threads
	result := DB.Conn.Table("threads").Select("count(id) as Q_Thread").
		Find(&Thread)

	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	return Thread.ToDomain(), nil
}

func (DB *MysqlThreadRepository) GetThreads(ctx context.Context) ([]threads.Domain, error) {
	var Thread []Threads
	result := DB.Conn.Table("threads").Select("threads.id, name, photo_url as Photo, title, category, threads.created_at").
		Joins("join users on threads.user_id = users.id").Joins("join categories on threads.category_id = categories.id").
		Order("threads.created_at desc").Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetHomepageThreads(ctx context.Context, id int) ([]threads.Domain, error) {
	var Thread []Threads
	var Comment comments.Comments

	Q_Like := DB.Conn.Table("thread_likes").Where("thread_id = threads.id").Select("count(thread_likes.user_id)").Group("thread_id")
	Q_Comment := DB.Conn.Table("comments").Where("thread_id = threads.id").Select("count(comment)").Group("thread_id")

	result := DB.Conn.Table("threads").Select("*, threads.id, title, content, (?) as Q_Like, (?) as Q_Comment", Q_Like, Q_Comment).
		Where("thread_follows.user_id = (?)", id).
		Joins("join users on threads.user_id = users.id").
		Joins("join thread_follows on threads.id = thread_follows.thread_id").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Table("comments").Select("*").
				Joins("join users on comments.user_id = users.id").Find(&Comment)
		}).
		// NESTED Comment but so hard to implement with GORM
		// Preload("Comments", func(db *gorm.DB) *gorm.DB {
		// 	return db.Table("comments").Select("*").Preload("Replies").Find(&Comment)
		// }).
		Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}
	fmt.Println(ToListDomain(Thread))
	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetRecommendationThreads(ctx context.Context, id int) ([]threads.Domain, error) {
	var Thread []Threads
	var Comment comments.Comments

	Q_Like := DB.Conn.Table("thread_likes").Where("thread_id = threads.id").Select("count(thread_likes.user_id)").Group("thread_id")
	Q_Comment := DB.Conn.Table("comments").Where("thread_id = threads.id").Select("count(comment)").Group("thread_id")

	result := DB.Conn.Table("threads").Select("*, threads.id, title, content, (?) as Q_Like, (?) as Q_Comment", Q_Like, Q_Comment).
		Where("threads.category_id = (?)", DB.GetCategories(id, 0)).
		Or("threads.category_id = (?)", DB.GetCategories(id, 1)).
		Or("threads.category_id = (?)", DB.GetCategories(id, 2)).
		Joins("join users on threads.user_id = users.id").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Table("comments").Select("*").
				Joins("join users on comments.user_id = users.id").Find(&Comment)
		}).
		Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetHotThreads(ctx context.Context) ([]threads.Domain, error) {
	var Thread []Threads
	var Comment comments.Comments

	Q_Like := DB.Conn.Table("thread_likes").Where("thread_id = threads.id").Select("count(thread_likes.user_id)").Group("thread_id")
	Q_Comment := DB.Conn.Table("comments").Where("thread_id = threads.id").Select("count(comment)").Group("thread_id")

	result := DB.Conn.Table("threads").Select("*, threads.id, title, content, (?) as Q_Like, (?) as Q_Comment", Q_Like, Q_Comment).
		Order("Q_Like desc, Q_Comment desc").
		Joins("join users on threads.user_id = users.id").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Table("comments").Select("*").
				Joins("join users on comments.user_id = users.id").Find(&Comment)
		}).
		Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) DeleteThread(ctx context.Context, id int) (threads.Domain, error) {
	var Thread Threads
	result := DB.Conn.Model(&Thread).Where("threads.id = ?", id).Update("active", "false")

	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	return Thread.ToDomain(), nil
}

func (DB *MysqlThreadRepository) GetCategories(id int, index int) *gorm.DB {
	Categories1 := DB.Conn.Table("threads").Select("categories.id").Where("users.id = (?)", id).Group("categories.id").
		Joins("join categories on threads.category_id = categories.id").
		Joins("join users on threads.user_id = users.id").
		Limit(1).Offset(index)

	return Categories1
}

func (DB *MysqlThreadRepository) GetSearch(ctx context.Context, threadname string) ([]threads.Domain, error) {
	var Thread []Threads
	var Comment comments.Comments
	var Newthreadname = ("%" + threadname + "%")

	Q_Like := DB.Conn.Table("thread_likes").Where("thread_id = threads.id").Select("count(thread_likes.user_id)").Group("thread_id")
	Q_Comment := DB.Conn.Table("comments").Where("thread_id = threads.id").Select("count(comment)").Group("thread_id")

	result := DB.Conn.Table("threads").Select("*, threads.id, title, content, (?) as Q_Like, (?) as Q_Comment", Q_Like, Q_Comment).
		Where("title LIKE ?", Newthreadname).
		Joins("join users on threads.user_id = users.id").
		Preload("Comments", func(db *gorm.DB) *gorm.DB {
			return db.Table("comments").Select("*, comments.id").
				Joins("join users on comments.user_id = users.id").Find(&Comment)
		}).
		Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) CreateThread(ctx context.Context, domain threads.Domain) (threads.Domain, error) {
	var Thread Threads

	result := DB.Conn.Model(&Thread).Create(&domain)

	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	return Thread.ToDomain(), nil
}
