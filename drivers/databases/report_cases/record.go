package reportcases

import (
	reportcases "fgd-alterra-29/business/report_cases"
	commentreport "fgd-alterra-29/drivers/databases/comment_report"
	threadreport "fgd-alterra-29/drivers/databases/thread_report"
)

type ReportCases struct {
	ID            int    `gorm:"primaryKey"`
	Case          string `gorm:"not null"`
	Description   string
	ThreadReport  []threadreport.ThreadReport   `gorm:"foreignKey:ReportCase_id"`
	CommentReport []commentreport.CommentReport `gorm:"foreignKey:ReportCase_id"`
}

func (Rc *ReportCases) ToDomain() reportcases.Domain {
	return reportcases.Domain{
		ID:          Rc.ID,
		Case:        Rc.Case,
		Description: Rc.Description,
	}
}

func ToListDomain(u []ReportCases) []reportcases.Domain {
	var Domains []reportcases.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
