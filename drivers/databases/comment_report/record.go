package commentreport

import (
	commentreport "fgd-alterra-29/business/comment_report"
	"time"
)

type CommentReport struct {
	ID            int `gorm:"primaryKey"`
	Comment_id    int
	Reporter_id   int
	ReportCase_id int
	Message       string
	Created_at    time.Time
	Q_Case        int    `gorm:"-:migration;->"`
	Case          string `gorm:"-:migration;->"`
	Comment       string `gorm:"-:migration;->"`
}

func (CR *CommentReport) ToDomain() commentreport.Domain {
	return commentreport.Domain{
		ID:          CR.ID,
		Comment_id:  CR.Comment_id,
		Reporter_id: CR.Reporter_id,
		Message:     CR.Message,
		Created_at:  CR.Created_at,
		Q_Case:      CR.Q_Case,
		Case:        CR.Case,
		Comment:     CR.Comment,
	}
}

func ToListDomain(u []CommentReport) []commentreport.Domain {
	var Domains []commentreport.Domain

	for _, val := range u {
		Domains = append(Domains, val.ToDomain())
	}
	return Domains
}
