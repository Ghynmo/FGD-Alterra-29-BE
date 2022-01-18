package routes

import (
	apinews "fgd-alterra-29/controllers/api_news"
	"fgd-alterra-29/controllers/badges"
	"fgd-alterra-29/controllers/categories"
	commentlikes "fgd-alterra-29/controllers/comment_likes"
	commentreport "fgd-alterra-29/controllers/comment_report"
	"fgd-alterra-29/controllers/comments"
	"fgd-alterra-29/controllers/follows"
	reportcases "fgd-alterra-29/controllers/report_cases"
	"fgd-alterra-29/controllers/reputations"
	threadlikes "fgd-alterra-29/controllers/thread_likes"
	threadreport "fgd-alterra-29/controllers/thread_report"
	threadsaves "fgd-alterra-29/controllers/thread_saves"
	threadshares "fgd-alterra-29/controllers/thread_shares"
	"fgd-alterra-29/controllers/threads"
	userbadges "fgd-alterra-29/controllers/user_badges"
	userpoints "fgd-alterra-29/controllers/user_points"
	"fgd-alterra-29/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JwtConfig               middleware.JWTConfig
	APINewsController       apinews.APINewsController
	UserController          users.UserController
	UserBadgeController     userbadges.UserBadgeController
	ThreadController        threads.ThreadController
	CommentController       comments.CommentController
	FollowController        follows.FollowController
	CategoryController      categories.CategoryController
	ReportCaseController    reportcases.ReportCaseController
	ThreadReportController  threadreport.ThreadReportController
	CommentReportController commentreport.CommentReportController
	ThreadLikeController    threadlikes.ThreadLikeController
	CommentLikeController   commentlikes.CommentLikeController
	ThreadSaveController    threadsaves.ThreadSaveController
	ThreadShareController   threadshares.ThreadShareController
	UserPointController     userpoints.UserPointController
	BadgeController         badges.BadgeController
	ReputationController    reputations.ReputationController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.Use(middleware.Logger())
	jwtAuth := middleware.JWTWithConfig(cl.JwtConfig)

	e.POST("register", cl.UserController.RegisterController)
	e.POST("login", cl.UserController.LoginController)

	//Profile Page
	e.GET("profile/:id", cl.UserController.GetProfileController, jwtAuth)
	e.GET("post/:id", cl.CommentController.GetProfileCommentsController, jwtAuth)
	e.GET("thread/:id", cl.ThreadController.GetProfileThreads, jwtAuth)
	e.GET("followers", cl.FollowController.GetFollowers)
	e.GET("following", cl.FollowController.GetFollowing)

	//Main Page of Admin
	e.GET("admin/", cl.UserController.GetDashboardController, jwtAuth)
	//List of Users (Admin Access)
	e.GET("admin/users", cl.UserController.GetUsersController, jwtAuth)
	e.GET("admin/users/search/:name", cl.UserController.GetUsersByName, jwtAuth)
	e.PUT("admin/users/banned/:id", cl.UserController.BannedUser, jwtAuth)
	//List of Threads (Admin Access)
	e.GET("admin/threads", cl.ThreadController.GetThreadsController, jwtAuth)
	e.GET("admin/threads/search/:title", cl.ThreadController.GetThreadsByTitleController, jwtAuth)
	e.PUT("admin/threads/thread/:id", cl.ThreadController.DeleteThread, jwtAuth)
	e.PUT("admin/threads/thread/:id", cl.ThreadController.ActivateThread, jwtAuth)
	//List of Posts (Admin Access)
	e.GET("admin/posts", cl.CommentController.GetPostsController, jwtAuth)
	e.GET("admin/posts/search/:comment", cl.CommentController.GetPostsByCommentController, jwtAuth)
	e.PUT("admin/posts/post/:id", cl.CommentController.UnactivatingPostController, jwtAuth)
	e.PUT("admin/posts/post/:id", cl.CommentController.ActivatingPostController, jwtAuth)
	//List of Comment's Reports (Admin Access)
	// e.GET("admin/comment-reports", cl.CommentReportController.AdminGetReports, jwtAuth)
	// e.GET("admin/comment-reports/search/:category", cl.CommentReportController.GetReportsByCategoryController, jwtAuth)
	// e.DELETE("admin/comment-reports/comment-report/:id", cl.CommentReportController.DeleteCommentReport, jwtAuth)
	//List of Thread's Reports (Admin Access)
	e.GET("admin/thread-reports", cl.ThreadReportController.AdminGetReports, jwtAuth)
	e.GET("admin/thread-reports/search/:category", cl.ThreadReportController.SearchReportsByCategoryController, jwtAuth)
	e.PUT("admin/thread-reports/thread-report/:id", cl.ThreadReportController.SolvedThreadReport, jwtAuth)
	//Edit Profile page
	e.GET("admin/edit/:id", cl.UserController.GetAdminSettingController, jwtAuth)
	e.GET("user/edit/:id", cl.UserController.GetUserSettingController, jwtAuth)
	e.PUT("admin/edit/:id", cl.UserController.UpdateAdminProfile, jwtAuth)
	e.PUT("user/edit/:id", cl.UserController.UpdateUserProfile, jwtAuth)

	e.GET("admin/report-thread", cl.ReportCaseController.GetReportForm, jwtAuth)
	e.POST("admin/report-thread", cl.ThreadReportController.CreateReportThread, jwtAuth)
	//Report Comment page
	// e.GET("admin/report-comment", cl.ReportCaseController.GetReportForm, jwtAuth)
	// e.POST("admin/report-comment", cl.CommentReportController.CreateReportComment, jwtAuth)
	//Report Thread page

	e.GET("home/:id", cl.ThreadController.GetHomepageThreads, jwtAuth)
	e.GET("recommendation/:id", cl.ThreadController.GetRecommendationThreads, jwtAuth)
	e.GET("hotthread", cl.ThreadController.GetHotThreads)
	e.GET("search", cl.ThreadController.GetSearch)
	e.GET("sidenews", cl.APINewsController.GetAPINewsController)

	e.GET("commentbythread", cl.CommentController.GetCommentByThreadController, jwtAuth)
	e.GET("comment/reply/:id", cl.CommentController.GetReplyComments, jwtAuth)
	e.POST("comment", cl.CommentController.CreateCommentController, jwtAuth)
	e.POST("thread", cl.ThreadController.CreateThread, jwtAuth, jwtAuth)
	e.POST("follows", cl.FollowController.FollowsController, jwtAuth)

	e.POST("commentlike", cl.CommentLikeController.Likes, jwtAuth)
	e.POST("threadlike", cl.ThreadLikeController.Likes, jwtAuth)
	e.POST("threadsave", cl.ThreadSaveController.SaveThread, jwtAuth)
	e.POST("threadshare", cl.ThreadShareController.ShareThread, jwtAuth)

	//additional
	e.POST("category", cl.CategoryController.CreateCategoryController, jwtAuth)
	e.POST("reportcase", cl.ReportCaseController.CreateCaseController, jwtAuth)
	e.POST("reputation", cl.ReputationController.CreateReputationController, jwtAuth)
}
