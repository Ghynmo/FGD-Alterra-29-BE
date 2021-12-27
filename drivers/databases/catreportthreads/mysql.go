package catreportthreads

import (
	"context"
	"fgd-alterra-29/business/catreportthreads"

	"gorm.io/gorm"
)

type MysqlCatReportThreadRepository struct {
	Conn *gorm.DB
}

func NewMysqlCatReportThreadRepository(conn *gorm.DB) catreportthreads.Repository {
	return &MysqlCatReportThreadRepository{
		Conn: conn,
	}
}

func (DB *MysqlCatReportThreadRepository) GetReportForm(ctx context.Context) ([]catreportthreads.Domain, error) {
	var catReportThread []CatReportT

	result := DB.Conn.Table("cat_report_ts").Select("id, category_report").Find(&catReportThread)

	if result.Error != nil {
		return []catreportthreads.Domain{}, result.Error
	}

	return ToListDomain(catReportThread), nil
}
