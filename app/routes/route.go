package routes

import (
	"fgd-alterra-29/app/middlewares"
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
	admin := e.Group("admin")
	profile := e.Group("profile")
	admin.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) { return middlewares.ExtractAdmin(c) }))

	//Main Page of Admin
	admin.GET("", cl.UserController.GetDashboardController, jwtAuth)

	//List of Users (Admin Access)
	admin.GET("/users", cl.UserController.GetUsersController, jwtAuth)
	admin.GET("/users/search/:name", cl.UserController.GetUsersByName, jwtAuth)
	admin.PUT("/users/banned/:id", cl.UserController.BannedUser, jwtAuth)

	//List of Threads (Admin Access)
	admin.GET("/threads", cl.ThreadController.GetThreadsController, jwtAuth)
	admin.GET("/threads/search/:title", cl.ThreadController.GetThreadsByTitleController, jwtAuth)
	admin.PUT("/threads/unactivate/:id", cl.ThreadController.DeleteThread, jwtAuth)
	admin.PUT("/threads/activate/:id", cl.ThreadController.ActivateThread, jwtAuth)

	//List of Posts (Admin Access)
	admin.GET("/posts", cl.CommentController.GetPostsController, jwtAuth)
	admin.GET("/posts/search/:comment", cl.CommentController.GetPostsByCommentController, jwtAuth)
	admin.PUT("/posts/unactivate/:id", cl.CommentController.UnactivatingPostController, jwtAuth)
	admin.PUT("/posts/activate/:id", cl.CommentController.ActivatingPostController, jwtAuth)

	//List of Comment's Reports (Admin Access)
	// admin.GET("/comment-reports", cl.CommentReportController.AdminGetReports, jwtAuth)
	// admin.GET("/comment-reports/search/:category", cl.CommentReportController.GetReportsByCategoryController, jwtAuth)
	// admin.DELETE("/comment-reports/comment-report/:id", cl.CommentReportController.DeleteCommentReport, jwtAuth)

	//List of Thread's Reports (Admin Access)
	admin.GET("/thread-reports", cl.ThreadReportController.AdminGetReports, jwtAuth)
	admin.GET("/thread-reports/search/:category", cl.ThreadReportController.SearchReportsByCategoryController, jwtAuth)
	admin.PUT("/thread-reports/thread-report/:id", cl.ThreadReportController.SolvedThreadReport, jwtAuth)

	//Edit Profile page
	admin.GET("/edit", cl.UserController.GetAdminSettingController, jwtAuth)
	admin.PUT("/edit", cl.UserController.UpdateAdminProfile, jwtAuth)

	//Report Comment page
	// admin.GET("/report-comment", cl.ReportCaseController.GetReportForm, jwtAuth)
	// admin.POST("/report-comment", cl.CommentReportController.CreateReportComment, jwtAuth)

	//additional
	admin.POST("/category", cl.CategoryController.CreateCategoryController, jwtAuth)
	admin.POST("/reportcase", cl.ReportCaseController.CreateCaseController, jwtAuth)
	admin.POST("/reputation", cl.ReputationController.CreateReputationController, jwtAuth)
	admin.POST("/badge", cl.BadgeController.CreateBadgeController, jwtAuth)
	// admin.PUT("/activate", cl.BadgeController.CreateBadgeController, jwtAuth)

	e.POST("register", cl.UserController.RegisterController)
	e.POST("login", cl.UserController.LoginController)

	//Profile Page
	profile.GET("/:id", cl.UserController.GetProfileController, jwtAuth)
	profile.GET("/post/:id", cl.CommentController.GetProfileCommentsController, jwtAuth)
	profile.GET("/thread/:id", cl.ThreadController.GetProfileThreads, jwtAuth)
	profile.GET("/followers", cl.FollowController.GetFollowers)
	profile.GET("/following", cl.FollowController.GetFollowing)
	profile.GET("/user/edit", cl.UserController.GetUserSettingController, jwtAuth)
	profile.PUT("/user/edit", cl.UserController.UpdateUserProfile, jwtAuth)

	e.GET("", cl.ThreadController.GetHotThreads)
	e.GET("home", cl.ThreadController.GetHomepageThreads, jwtAuth)
	e.GET("recommendation", cl.ThreadController.GetRecommendationThreads, jwtAuth)
	e.GET("hotthread", cl.ThreadController.GetHotThreads)
	e.GET("search", cl.ThreadController.GetSearch)
	e.GET("thread/:id", cl.ThreadController.GetThreadsByIDController, jwtAuth)
	e.GET("sidenews", cl.APINewsController.GetAPINewsController)

	e.GET("categories", cl.CategoryController.GetCategoriesController, jwtAuth)

	e.GET("commentbythread", cl.CommentController.GetCommentByThreadController, jwtAuth)
	e.GET("comment/reply/:reply_of", cl.CommentController.GetReplyComments, jwtAuth)
	e.POST("comment", cl.CommentController.CreateCommentController, jwtAuth)
	e.POST("thread", cl.ThreadController.CreateThread, jwtAuth, jwtAuth)
	e.POST("follows", cl.FollowController.FollowsController, jwtAuth)

	e.POST("commentlike", cl.CommentLikeController.Likes, jwtAuth)
	e.POST("threadlike", cl.ThreadLikeController.Likes, jwtAuth)
	e.POST("threadsave", cl.ThreadSaveController.SaveThread, jwtAuth)
	e.POST("threadshare", cl.ThreadShareController.ShareThread, jwtAuth)

	e.GET("report-thread", cl.ReportCaseController.GetReportForm, jwtAuth)
	e.POST("report-thread", cl.ThreadReportController.CreateReportThread, jwtAuth)

}
