package threadreport

import (
	"context"
	threadreport "fgd-alterra-29/business/thread_report"

	"gorm.io/gorm"
)

type MysqlThreadReportRepository struct {
	Conn *gorm.DB
}

func NewMysqlThreadReportRepository(conn *gorm.DB) threadreport.Repository {
	return &MysqlThreadReportRepository{
		Conn: conn,
	}
}

func (DB *MysqlThreadReportRepository) GetReportsByCategory(ctx context.Context, category string) ([]threadreport.Domain, error) {
	var ThreadReport []ThreadReport
	var NewCategory = ("%" + category + "%")

	result := DB.Conn.Table("thread_reports").Where("category_report LIKE ?", NewCategory).Select("thread_reports.id, title as Thread, message, category_report as CategoryReport").
		Joins("join cat_report_ts on thread_reports.catreportthread_id = cat_report_ts.id").
		Joins("join threads on thread_reports.thread_id = threads.id").Order("thread_reports.created_at").
		Find(&ThreadReport)

	if result.Error != nil {
		return []threadreport.Domain{}, result.Error
	}

	return ToListDomain(ThreadReport), nil
}

func (DB *MysqlThreadReportRepository) GetThreadReports(ctx context.Context) ([]threadreport.Domain, error) {
	var ThreadReport []ThreadReport
	result := DB.Conn.Table("thread_reports").Select("thread_reports.id ,count(thread_reports.id) as Q_Cat, category_report as CategoryReport").
		Joins("join cat_report_ts on thread_reports.catreportthread_id = cat_report_ts.id").Group("category_report").
		Find(&ThreadReport)

	if result.Error != nil {
		return []threadreport.Domain{}, result.Error
	}

	return ToListDomain(ThreadReport), nil
}

func (DB *MysqlThreadReportRepository) CreateReportThread(ctx context.Context, domain threadreport.Domain) (threadreport.Domain, error) {
	var ThreadReport ThreadReport

	result := DB.Conn.Model(&ThreadReport).Create(&domain)

	if result.Error != nil {
		return threadreport.Domain{}, result.Error
	}

	return ThreadReport.ToDomain(), nil
}

func (DB *MysqlThreadReportRepository) GetReports(ctx context.Context) ([]threadreport.Domain, error) {
	var ThreadReport []ThreadReport
	result := DB.Conn.Table("thread_reports").Select("thread_reports.id, title as Thread, message, category_report as CategoryReport").
		Joins("join cat_report_ts on thread_reports.catreportthread_id = cat_report_ts.id").
		Joins("join threads on thread_reports.thread_id = threads.id").Order("thread_reports.created_at").
		Find(&ThreadReport)

	if result.Error != nil {
		return []threadreport.Domain{}, result.Error
	}

	return ToListDomain(ThreadReport), nil
}

func (DB *MysqlThreadReportRepository) DeleteThreadReport(ctx context.Context, id int) (threadreport.Domain, error) {
	var ThreadReport ThreadReport
	result := DB.Conn.Delete(&ThreadReport, id)

	if result.Error != nil {
		return threadreport.Domain{}, result.Error
	}

	return ThreadReport.ToDomain(), nil
}
