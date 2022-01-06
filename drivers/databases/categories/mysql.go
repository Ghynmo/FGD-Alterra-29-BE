package categories

import (
	"context"
	"fgd-alterra-29/business/categories"

	"gorm.io/gorm"
)

type MysqlCategoryRepository struct {
	Conn *gorm.DB
}

func NewMysqlCategoryRepository(conn *gorm.DB) categories.Repository {
	return &MysqlCategoryRepository{
		Conn: conn,
	}
}

func (DB *MysqlCategoryRepository) GetUserActiveInCategory(ctx context.Context, id int) ([]categories.Domain, error) {
	var Category []Categories

	result := DB.Conn.Table("categories").Select("count(title) as Q_Title, category, icon_url").Where("threads.user_id = 1").
		Joins("join threads on categories.id = threads.category_id").
		Group("category_id").
		Find(&Category)

	if result.Error != nil {
		return []categories.Domain{}, result.Error
	}

	return ToListDomain(Category), nil
}
