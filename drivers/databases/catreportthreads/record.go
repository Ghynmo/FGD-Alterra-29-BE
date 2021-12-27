package catreportthreads

import (
	"fgd-alterra-29/business/catreportthreads"
	threadreport "fgd-alterra-29/drivers/databases/thread_report"
)

type CatReportT struct {
	ID             int    `gorm:"primaryKey"`
	CategoryReport string `gorm:"not null"`
	Description    string
	ThreadReport   []threadreport.ThreadReport `gorm:"foreignKey:Catreportthread_id"`
}

func (Crt *CatReportT) ToDomain() catreportthreads.Domain {
	return catreportthreads.Domain{
		ID:             Crt.ID,
		CategoryReport: Crt.CategoryReport,
		Description:    Crt.Description,
	}
}

func ToListDomain(u []CatReportT) []catreportthreads.Domain {
	var Domains []catreportthreads.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
