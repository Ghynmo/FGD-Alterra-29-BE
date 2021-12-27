package responses

import (
	"fgd-alterra-29/business/comments"
	threadreport "fgd-alterra-29/business/thread_report"
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/business/users"
	"fgd-alterra-29/controllers/users/responses/dashboard"
)

type Dashboard struct {
	Users      []dashboard.UserList
	ThreadStat []dashboard.ThreadReportStat
	Q_Users    int
	Q_Thread   int
	Q_Post     int
}

func ToDashboard(users []users.Domain, treport []threadreport.Domain, quser users.Domain, qthread threads.Domain, qpost comments.Domain) Dashboard {
	return Dashboard{
		Users:      dashboard.ToListUserList(users),
		ThreadStat: dashboard.ToListThreadReportStat(treport),
		Q_Users:    quser.Q_User,
		Q_Thread:   qthread.Q_Thread,
		Q_Post:     qpost.Q_Post,
	}
}
