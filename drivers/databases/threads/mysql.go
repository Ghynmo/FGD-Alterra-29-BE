package threads

import (
	"context"
	"fgd-alterra-29/business/threads"
	"time"

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
		Where("title LIKE ? AND threads.active = 1", NewTitle).Order("threads.created_at desc").Find(&Thread)

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

	result := DB.Conn.Table("threads").Where("threads.user_id = ? AND threads.active = 1", id).Select("id, title, (?) as Comment, (?) as RecentReplier", Comment, Replier).
		Order("created_at desc").Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetThreadQuantity(ctx context.Context) (threads.Domain, error) {
	var Thread Threads
	result := DB.Conn.Table("threads").Select("count(id) as Q_Thread").Where("threads.active = 1").
		Find(&Thread)

	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	return Thread.ToDomain(), nil
}

func (DB *MysqlThreadRepository) GetThreadQtyByCategory(ctx context.Context, domain threads.Domain, id int) (threads.Domain, error) {

	var Thread Threads
	result := DB.Conn.Table("threads").Select("count(title) as Q_Thread").
		Where("threads.active = 1 AND category_id = ? AND user_id = ?", domain.Category_id, id).
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

	Q_Like := DB.Conn.Table("thread_likes").Where("thread_id = threads.id").Select("count(thread_likes.user_id)").Group("thread_id")
	Q_Comment := DB.Conn.Table("comments").Where("thread_id = threads.id").Select("count(comment)").Group("thread_id")

	result := DB.Conn.Table("threads").Select("*, threads.id, title, content, (?) as Q_Like, (?) as Q_Comment", Q_Like, Q_Comment).
		Where("follows.follower_id = (?) AND threads.active = 1", id).Order("threads.created_at desc").
		Joins("join users on threads.user_id = users.id").
		Joins("join follows on users.id = follows.user_id").
		Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}
	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetRecommendationThreads(ctx context.Context, id int) ([]threads.Domain, error) {
	var Thread []Threads

	Q_Like := DB.Conn.Table("thread_likes").Where("thread_id = threads.id").Select("count(thread_likes.user_id)").Group("thread_id")
	Q_Comment := DB.Conn.Table("comments").Where("thread_id = threads.id").Select("count(comment)").Group("thread_id")

	result := DB.Conn.Table("threads").Select("*, threads.id, title, content, (?) as Q_Like, (?) as Q_Comment", Q_Like, Q_Comment).
		Where("threads.category_id = (?) AND threads.active = 1", DB.GetCategories(id, 0)).
		Or("threads.category_id = (?)", DB.GetCategories(id, 1)).
		Or("threads.category_id = (?)", DB.GetCategories(id, 2)).
		Joins("join users on threads.user_id = users.id").
		Order("threads.created_at desc").
		Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetHotThreads(ctx context.Context) ([]threads.Domain, error) {
	var Thread []Threads

	Q_Like := DB.Conn.Table("thread_likes").Where("thread_id = threads.id").Select("count(thread_likes.user_id)").Group("thread_id")
	Q_Comment := DB.Conn.Table("comments").Where("thread_id = threads.id").Select("count(comment)").Group("thread_id")

	result := DB.Conn.Table("threads").Select("*, threads.id, title, content, (?) as Q_Like, (?) as Q_Comment", Q_Like, Q_Comment).
		Where("threads.active = 1").Order("Q_Like desc, Q_Comment desc").
		Joins("join users on threads.user_id = users.id").
		Order("threads.created_at desc").
		Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetThreadByID(ctx context.Context, id int) (threads.Domain, error) {
	var Thread Threads

	Q_Like := DB.Conn.Table("thread_likes").Where("thread_id = threads.id").Select("count(thread_likes.user_id)").Group("thread_id")
	Q_Comment := DB.Conn.Table("comments").Where("thread_id = threads.id").Select("count(comment)").Group("thread_id")

	result := DB.Conn.Table("threads").Select("*, threads.id, title, content, (?) as Q_Like, (?) as Q_Comment", Q_Like, Q_Comment).
		Where("threads.id = ?", id).
		Joins("join users on threads.user_id = users.id").
		Find(&Thread)

	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	return Thread.ToDomain(), nil
}

func (DB *MysqlThreadRepository) DeleteThread(ctx context.Context, id int) (threads.Domain, error) {
	var Thread Threads
	result := DB.Conn.Model(&Thread).Where("threads.id = ?", id).Update("active", 0)

	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	return Thread.ToDomain(), nil
}

func (DB *MysqlThreadRepository) ActivateThread(ctx context.Context, id int) (threads.Domain, error) {
	var Thread Threads
	result := DB.Conn.Model(&Thread).Where("threads.id = ?", id).Update("active", 1)

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
	var Newthreadname = ("%" + threadname + "%")

	Q_Like := DB.Conn.Table("thread_likes").Where("thread_id = threads.id").Select("count(thread_likes.user_id)").Group("thread_id")
	Q_Comment := DB.Conn.Table("comments").Where("thread_id = threads.id").Select("count(comment)").Group("thread_id")

	result := DB.Conn.Table("threads").Select("*, threads.id, title, content, (?) as Q_Like, (?) as Q_Comment", Q_Like, Q_Comment).
		Where("title LIKE ? AND threads.active = 1", Newthreadname).
		Joins("join users on threads.user_id = users.id").
		Order("threads.created_at desc").
		Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetThreadsByCategoryID(ctx context.Context, id int) ([]threads.Domain, error) {
	var Thread []Threads

	Q_Like := DB.Conn.Table("thread_likes").Where("thread_id = threads.id").Select("count(thread_likes.user_id)").Group("thread_id")
	Q_Comment := DB.Conn.Table("comments").Where("thread_id = threads.id").Select("count(comment)").Group("thread_id")

	result := DB.Conn.Table("threads").Select("*, threads.id, title, content, (?) as Q_Like, (?) as Q_Comment", Q_Like, Q_Comment).
		Where("threads.category_id = ? AND threads.active = 1", id).
		Joins("join users on threads.user_id = users.id").
		Order("threads.created_at desc").
		Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) GetSideNewsThreads(ctx context.Context) ([]threads.Domain, error) {
	var Thread []Threads

	result := DB.Conn.Table("threads").Select("threads.id, title, thumbnail_url").Where("threads.active = 1").
		Order("threads.created_at desc").Find(&Thread)

	if result.Error != nil {
		return []threads.Domain{}, result.Error
	}

	return ToListDomain(Thread), nil
}

func (DB *MysqlThreadRepository) CreateThread(ctx context.Context, domain threads.Domain, id int) (threads.Domain, error) {

	data := Threads{
		User_id:     id,
		Title:       domain.Title,
		Category_id: domain.Category_id,
		Content:     domain.Content,
		Created_at:  time.Now(),
	}

	result := DB.Conn.Model(&data).Create(&data)

	if result.Error != nil {
		return threads.Domain{}, result.Error
	}

	return data.ToDomain(), nil
}
