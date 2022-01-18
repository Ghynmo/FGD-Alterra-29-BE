package threadreport

import (
	threadreport "fgd-alterra-29/business/thread_report"
	"time"
)

type ThreadReport struct {
	ID            int `gorm:"primaryKey"`
	Thread_id     int
	Reporter_id   int
	ReportCase_id int
	Message       string
	State         string `gorm:"default:pending"`
	Created_at    time.Time
	Updated_at    time.Time
	Deleted_at    time.Time
	Q_Case        int    `gorm:"-:migration;->"`
	Case          string `gorm:"-:migration;->"`
	Thread        string `gorm:"-:migration;->"`
}

func (TR *ThreadReport) ToDomain() threadreport.Domain {
	return threadreport.Domain{
		ID:            TR.ID,
		Thread_id:     TR.Thread_id,
		Reporter_id:   TR.Reporter_id,
		ReportCase_id: TR.ReportCase_id,
		Message:       TR.Message,
		State:         TR.State,
		Created_at:    TR.Created_at,
		Updated_at:    TR.Updated_at,
		Deleted_at:    TR.Deleted_at,
		Q_Case:        TR.Q_Case,
		Case:          TR.Case,
		Thread:        TR.Thread,
	}
}

func ToListDomain(u []ThreadReport) []threadreport.Domain {
	var Domains []threadreport.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
