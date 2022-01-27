package threadreport

import (
	"context"
	threadreport "fgd-alterra-29/business/thread_report"
	"time"

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

func (DB *MysqlThreadReportRepository) SearchReportsByCategory(ctx context.Context, category string) ([]threadreport.Domain, error) {
	var ThreadReport []ThreadReport
	var NewCategory = ("%" + category + "%")

	result := DB.Conn.Table("thread_reports").Where("report_cases.case LIKE ?", NewCategory).Select("thread_reports.id, title as Thread, message, report_cases.case").
		Joins("join report_cases on thread_reports.report_case_id = report_cases.id").
		Joins("join threads on thread_reports.thread_id = threads.id").Order("thread_reports.created_at").
		Find(&ThreadReport)

	if result.Error != nil {
		return []threadreport.Domain{}, result.Error
	}

	return ToListDomain(ThreadReport), nil
}

func (DB *MysqlThreadReportRepository) GetThreadReportStat(ctx context.Context) ([]threadreport.Domain, error) {
	var ThreadReport []ThreadReport
	result := DB.Conn.Table("thread_reports").Select("thread_reports.id, count(thread_reports.id) as Q_Case, report_cases.case").
		Joins("join report_cases on thread_reports.report_case_id = report_cases.id").Group("report_cases.case").
		Find(&ThreadReport)

	if result.Error != nil {
		return []threadreport.Domain{}, result.Error
	}

	return ToListDomain(ThreadReport), nil
}

func (DB *MysqlThreadReportRepository) CreateReportThread(ctx context.Context, domain threadreport.Domain, my_id int) (threadreport.Domain, error) {

	data := ThreadReport{
		Reporter_id:   my_id,
		Thread_id:     domain.Thread_id,
		ReportCase_id: domain.ReportCase_id,
		Message:       domain.Message,
		Created_at:    time.Now(),
	}

	result := DB.Conn.Model(&data).Create(&data)

	if result.Error != nil {
		return threadreport.Domain{}, result.Error
	}

	return data.ToDomain(), nil
}

func (DB *MysqlThreadReportRepository) AdminGetReports(ctx context.Context) ([]threadreport.Domain, error) {
	var ThreadReport []ThreadReport
	result := DB.Conn.Table("thread_reports").Select("thread_reports.id, title as Thread, message, report_cases.case").
		Joins("join report_cases on thread_reports.report_case_id = report_cases.id").
		Joins("join threads on thread_reports.thread_id = threads.id").Order("thread_reports.created_at").
		Find(&ThreadReport)

	if result.Error != nil {
		return []threadreport.Domain{}, result.Error
	}

	return ToListDomain(ThreadReport), nil
}

func (DB *MysqlThreadReportRepository) SolvedThreadReport(ctx context.Context, id int) (threadreport.Domain, error) {
	var ThreadReport ThreadReport
	result := DB.Conn.Model(&ThreadReport).Where("thread_reports.id = ?", id).
		Update("status", "solved")
	if result.Error != nil {
		return threadreport.Domain{}, result.Error
	}

	return ThreadReport.ToDomain(), nil
}
