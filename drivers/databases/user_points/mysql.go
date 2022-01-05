package userpoints

import (
	"context"
	userpoint "fgd-alterra-29/business/user_points"
	"fmt"

	"gorm.io/gorm"
)

type MysqlUserPointRepository struct {
	Conn *gorm.DB
}

func NewMysqlUserPointRepository(conn *gorm.DB) userpoint.Repository {
	return &MysqlUserPointRepository{
		Conn: conn,
	}
}

func (DB *MysqlUserPointRepository) AddThreadPoint(ctx context.Context, id int) (userpoint.Domain, error) {
	var UserPoint UserPoints
	fmt.Println("uid", id)

	result := DB.Conn.Model(&UserPoint).Where("user_id = ?", id).Update("thread_point", gorm.Expr("thread_point + 10"))

	if result.Error != nil {
		return userpoint.Domain{}, result.Error
	}
	return userpoint.Domain{}, nil
}

func (DB *MysqlUserPointRepository) AddPostPoint(ctx context.Context, id int) (userpoint.Domain, error) {
	var UserPoint UserPoints

	result := DB.Conn.Table("users").Find(&UserPoint)

	if result.Error != nil {
		return userpoint.Domain{}, result.Error
	}
	return userpoint.Domain{}, nil
}
