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

	result := DB.Conn.Table("report_cases").Select("id, report_cases.case").Find(&reportcase)

	if result.Error != nil {
		return []reportcases.Domain{}, result.Error
	}

	return ToListDomain(reportcase), nil
}
