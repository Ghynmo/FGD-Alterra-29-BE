package responses

import (
	commentreport "fgd-alterra-29/business/comment_report"
	"fgd-alterra-29/business/comments"
	"fgd-alterra-29/business/threads"
	"fgd-alterra-29/business/users"
	"fgd-alterra-29/controllers/users/responses/dashboard"
)

type Dashboard struct {
	Users    []dashboard.UserList          `json:"users_list"`
	PostStat []dashboard.CommentReportStat `json:"post_report_statistic"`
	Q_Users  int                           `json:"users_total"`
	Q_Thread int                           `json:"threads_total"`
	Q_Post   int                           `json:"posts_total"`
}

func ToDashboard(users []users.Domain, quser users.Domain, qthread threads.Domain, qpost comments.Domain, creport []commentreport.Domain) Dashboard {
	return Dashboard{
		Users:    dashboard.ToListUserList(users),
		PostStat: dashboard.ToListCommentReportStat(creport),
		Q_Users:  quser.Q_User,
		Q_Thread: qthread.Q_Thread,
		Q_Post:   qpost.Q_Post,
	}
}
