package commentreport

import (
	"context"
	commentreport "fgd-alterra-29/business/comment_report"

	"gorm.io/gorm"
)

type MysqlCommentReportRepository struct {
	Conn *gorm.DB
}

func NewMysqlCommentReportRepository(conn *gorm.DB) commentreport.Repository {
	return &MysqlCommentReportRepository{
		Conn: conn,
	}
}

func (DB *MysqlCommentReportRepository) SearchReportsByCategory(ctx context.Context, reportcase string) ([]commentreport.Domain, error) {
	var CommentReport []CommentReport
	var NewReportCase = ("%" + reportcase + "%")

	result := DB.Conn.Table("comment_reports").Where("report_cases.case LIKE ?", NewReportCase).Select("comment_reports.id, comment, message, report_cases.case").
		Joins("join report_cases on comment_reports.report_case_id = report_cases.id").
		Joins("join comments on comment_reports.comment_id = comments.id").Order("comment_reports.created_at").
		Find(&CommentReport)

	if result.Error != nil {
		return []commentreport.Domain{}, result.Error
	}

	return ToListDomain(CommentReport), nil
}

func (DB *MysqlCommentReportRepository) GetCommentReportStat(ctx context.Context) ([]commentreport.Domain, error) {
	var CommentReport []CommentReport
	result := DB.Conn.Table("comment_reports").Select("comment_reports.id, count(comment_reports.id) as Q_Case, report_cases.case").
		Joins("join report_cases on comment_reports.report_case_id = report_cases.id").Group("report_cases.case").
		Find(&CommentReport)

	if result.Error != nil {
		return []commentreport.Domain{}, result.Error
	}

	return ToListDomain(CommentReport), nil
}

func (DB *MysqlCommentReportRepository) CreateReportComment(ctx context.Context, domain commentreport.Domain) (commentreport.Domain, error) {
	var CommentReport CommentReport

	result := DB.Conn.Model(&CommentReport).Create(&domain)

	if result.Error != nil {
		return commentreport.Domain{}, result.Error
	}

	return CommentReport.ToDomain(), nil
}

func (DB *MysqlCommentReportRepository) AdminGetReports(ctx context.Context) ([]commentreport.Domain, error) {
	var CommentReport []CommentReport
	result := DB.Conn.Table("comment_reports").Select("comment_reports.id, comment, message, report_cases.case").
		Joins("join report_cases on comment_reports.report_case_id = report_cases.id").
		Joins("join comments on comment_reports.comment_id = comments.id").Order("comment_reports.created_at").
		Find(&CommentReport)

	if result.Error != nil {
		return []commentreport.Domain{}, result.Error
	}

	return ToListDomain(CommentReport), nil
}

func (DB *MysqlCommentReportRepository) DeleteCommentReport(ctx context.Context, id int) (commentreport.Domain, error) {
	var CommentReport CommentReport
	result := DB.Conn.Delete(&CommentReport, id)

	if result.Error != nil {
		return commentreport.Domain{}, result.Error
	}

	return CommentReport.ToDomain(), nil
}
