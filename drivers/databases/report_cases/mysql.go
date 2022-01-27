package reportcases

import (
	"context"
	reportcases "fgd-alterra-29/business/report_cases"

	"gorm.io/gorm"
)

type MysqlReportCaseRepository struct {
	Conn *gorm.DB
}

func NewMysqlReportCaseRepository(conn *gorm.DB) reportcases.Repository {
	return &MysqlReportCaseRepository{
		Conn: conn,
	}
}

func (DB *MysqlReportCaseRepository) GetReportForm(ctx context.Context) ([]reportcases.Domain, error) {
	var reportcase []ReportCases

	result := DB.Conn.Table("report_cases").Select("id, report_cases.case").
		Order("report_cases.case asc").Find(&reportcase)

	if result.Error != nil {
		return []reportcases.Domain{}, result.Error
	}

	return ToListDomain(reportcase), nil
}

func (DB *MysqlReportCaseRepository) CreateCase(ctx context.Context, domain reportcases.Domain) (reportcases.Domain, error) {
	var reportcase ReportCases

	result := DB.Conn.Model(&reportcase).Create(&domain)

	if result.Error != nil {
		return reportcases.Domain{}, result.Error
	}

	return reportcase.ToDomain(), nil
}
