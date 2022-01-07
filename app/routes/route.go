package routes

import (
	"fgd-alterra-29/controllers/categories"
	"fgd-alterra-29/controllers/catreportthreads"
	"fgd-alterra-29/controllers/comments"
	"fgd-alterra-29/controllers/follows"
	threadreport "fgd-alterra-29/controllers/thread_report"
	"fgd-alterra-29/controllers/threads"
	userbadges "fgd-alterra-29/controllers/user_badges"
	"fgd-alterra-29/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController            users.UserController
	UserBadgeController       userbadges.UserBadgeController
	ThreadController          threads.ThreadController
	CommentController         comments.CommentController
	FollowController          follows.FollowController
	CategoryController        categories.CategoryController
	CatReportThreadController catreportthreads.CatReportThreadController
	ThreadReportController    threadreport.ThreadReportController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.GET("x", cl.UserController.GetUsersController)
	e.GET("profile/:id", cl.UserController.GetProfileController)
	e.GET("post/:id", cl.CommentController.GetProfileCommentsController)
	e.GET("thread/:id", cl.ThreadController.GetProfileThreads)
	e.GET("followers/:id", cl.FollowController.GetFollowers)
	e.GET("following/:id", cl.FollowController.GetFollowing)

	//Main Page of Admin
	e.GET("admin/", cl.UserController.GetDashboardController)

	//List of Users (Admin Access)
	e.GET("admin/users", cl.UserController.GetUsersController)
	e.GET("admin/users/search/:name", cl.UserController.GetUsersByName)
	e.PUT("admin/users/banned/:id", cl.UserController.BannedUser)
	e.PUT("admin/users/unbanned/:id", cl.UserController.UnbannedUser)

	//List of Threads (Admin Access)
	e.GET("admin/threads", cl.ThreadController.GetThreadsController)
	e.GET("admin/threads/search/:title", cl.ThreadController.GetThreadsByTitleController)
	e.DELETE("admin/threads/thread/:id", cl.ThreadController.DeleteThread)

	//List of Posts (Admin Access)
	e.GET("admin/posts", cl.CommentController.GetPostsController)
	e.GET("admin/posts/search/:comment", cl.CommentController.GetPostsByCommentController)
	e.DELETE("admin/posts/post/:id", cl.CommentController.UnactivatingPostController)

	//List of Reports (Admin Access)
	e.GET("admin/thread-reports", cl.ThreadReportController.GetReports)
	e.GET("admin/thread-reports/search/:category", cl.ThreadReportController.GetReportsByCategoryController)
	e.DELETE("admin/thread-reports/thread-report/:id", cl.ThreadReportController.DeleteThreadReport)

	//Edit Profile page
	e.GET("admin/edit/:id", cl.UserController.GetAdminSettingController)
	e.GET("user/edit/:id", cl.UserController.GetUserSettingController)
	e.PUT("admin/edit", cl.UserController.UpdateAdminProfile)
	e.PUT("user/edit", cl.UserController.UpdateUserProfile)

	//Report Thread page
	e.GET("admin/report-thread", cl.CatReportThreadController.GetReportForm)
	e.POST("admin/report-thread", cl.ThreadReportController.CreateReportThread)
}
