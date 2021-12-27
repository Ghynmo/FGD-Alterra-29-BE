package threadreport

import (
	threadreport "fgd-alterra-29/business/thread_report"
	"time"
)

type ThreadReport struct {
	ID                 int `gorm:"primaryKey"`
	Thread_id          int
	User_id            int
	Catreportthread_id int
	Message            string
	Created_at         time.Time
	Updated_at         time.Time
	Deleted_at         time.Time
	Q_Cat              int    `gorm:"-:migration;->"`
	CategoryReport     string `gorm:"-:migration;->"`
	Thread             string `gorm:"-:migration;->"`
}

func (TR *ThreadReport) ToDomain() threadreport.Domain {
	return threadreport.Domain{
		ID:             TR.ID,
		Thread_id:      TR.Thread_id,
		User_id:        TR.User_id,
		ReportGroup_id: TR.Catreportthread_id,
		Message:        TR.Message,
		Created_at:     TR.Created_at,
		Updated_at:     TR.Updated_at,
		Deleted_at:     TR.Deleted_at,
		Q_Cat:          TR.Q_Cat,
		CategoryReport: TR.CategoryReport,
		Thread:         TR.Thread,
	}
}

func ToListDomain(u []ThreadReport) []threadreport.Domain {
	var Domains []threadreport.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
