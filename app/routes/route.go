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
	"net/http"

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
	corsMid := middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type"},
		AllowCredentials: true,
	})
	// corsMid := cors.New(cors.Options{
	// 	AllowedOrigins: []string{"*"},
	// 	AllowedMethods: []string{"OPTIONS", "GET", "POST", "PUT"},
	// 	AllowedHeaders: []string{"*"},
	// 	Debug:          true,
	// })
	// e.Use(echo.WrapMiddleware(corsMid.Handler))
	jwtAuth := middleware.JWTWithConfig(cl.JwtConfig)

	e.POST("register", cl.UserController.RegisterController, corsMid)
	e.POST("login", cl.UserController.LoginController, corsMid)

	//Profile Page
	e.GET("profile/:id", cl.UserController.GetProfileController, jwtAuth, corsMid)
	e.GET("post/:id", cl.CommentController.GetProfileCommentsController, jwtAuth, corsMid)
	e.GET("thread/:id", cl.ThreadController.GetProfileThreads, jwtAuth, corsMid)
	e.GET("followers", cl.FollowController.GetFollowers, corsMid)
	e.GET("following", cl.FollowController.GetFollowing, corsMid)

	//Main Page of Admin
	e.GET("admin/", cl.UserController.GetDashboardController, jwtAuth, corsMid)
	//List of Users (Admin Access)
	e.GET("admin/users", cl.UserController.GetUsersController, jwtAuth, corsMid)
	e.GET("admin/users/search/:name", cl.UserController.GetUsersByName, jwtAuth, corsMid)
	e.PUT("admin/users/banned/:id", cl.UserController.BannedUser, jwtAuth, corsMid)
	e.PUT("admin/users/unbanned/:id", cl.UserController.UnbannedUser, jwtAuth, corsMid)
	//List of Threads (Admin Access)
	e.GET("admin/threads", cl.ThreadController.GetThreadsController, jwtAuth, corsMid)
	e.GET("admin/threads/search/:title", cl.ThreadController.GetThreadsByTitleController, jwtAuth, corsMid)
	e.DELETE("admin/threads/thread/:id", cl.ThreadController.DeleteThread, jwtAuth, corsMid)
	e.PUT("admin/threads/thread/:id", cl.ThreadController.ActivateThread, jwtAuth, corsMid)
	//List of Posts (Admin Access)
	e.GET("admin/posts", cl.CommentController.GetPostsController, jwtAuth, corsMid)
	e.GET("admin/posts/search/:comment", cl.CommentController.GetPostsByCommentController, jwtAuth, corsMid)
	e.DELETE("admin/posts/post/:id", cl.CommentController.UnactivatingPostController, jwtAuth, corsMid)
	e.PUT("admin/posts/post/:id", cl.CommentController.ActivatingPostController, jwtAuth, corsMid)
	//List of Comment's Reports (Admin Access)
	e.GET("admin/comment-reports", cl.CommentReportController.AdminGetReports, jwtAuth, corsMid)
	e.GET("admin/comment-reports/search/:category", cl.CommentReportController.GetReportsByCategoryController, jwtAuth, corsMid)
	e.DELETE("admin/comment-reports/comment-report/:id", cl.CommentReportController.DeleteCommentReport, jwtAuth, corsMid)
	//List of Thread's Reports (Admin Access)
	e.GET("admin/thread-reports", cl.ThreadReportController.AdminGetReports, jwtAuth, corsMid)
	e.GET("admin/thread-reports/search/:category", cl.ThreadReportController.SearchReportsByCategoryController, jwtAuth, corsMid)
	e.DELETE("admin/thread-reports/thread-report/:id", cl.ThreadReportController.DeleteThreadReport, jwtAuth, corsMid)
	//Edit Profile page
	e.GET("admin/edit/:id", cl.UserController.GetAdminSettingController, jwtAuth, corsMid)
	e.GET("user/edit/:id", cl.UserController.GetUserSettingController, jwtAuth, corsMid)
	e.PUT("admin/edit/:id", cl.UserController.UpdateAdminProfile, jwtAuth, corsMid)
	e.PUT("user/edit/:id", cl.UserController.UpdateUserProfile, jwtAuth, corsMid)

	e.GET("admin/report-thread", cl.ReportCaseController.GetReportForm, jwtAuth, corsMid)
	e.POST("admin/report-thread", cl.ThreadReportController.CreateReportThread, jwtAuth, corsMid)
	//Report Comment page
	// e.GET("admin/report-comment", cl.ReportCaseController.GetReportForm, jwtAuth, corsMid)
	// e.POST("admin/report-comment", cl.CommentReportController.CreateReportComment, jwtAuth, corsMid)
	//Report Thread page

	e.GET("home/:id", cl.ThreadController.GetHomepageThreads, jwtAuth, corsMid)
	e.GET("recommendation/:id", cl.ThreadController.GetRecommendationThreads, jwtAuth, corsMid)
	e.GET("hotthread", cl.ThreadController.GetHotThreads, corsMid)
	e.GET("search", cl.ThreadController.GetSearch, corsMid)
	e.GET("sidenews", cl.APINewsController.GetAPINewsController, corsMid)

	e.GET("commentbythread", cl.CommentController.GetCommentByThreadController, jwtAuth, corsMid)
	e.GET("comment/reply/:id", cl.CommentController.GetReplyComments, jwtAuth, corsMid)
	e.POST("comment", cl.CommentController.CreateCommentController, jwtAuth, corsMid)
	e.POST("thread", cl.ThreadController.CreateThread, jwtAuth, jwtAuth, corsMid)
	e.POST("follows", cl.FollowController.FollowsController, jwtAuth, corsMid)

	e.POST("commentlike", cl.CommentLikeController.Likes, jwtAuth, corsMid)
	e.POST("threadlike", cl.ThreadLikeController.Likes, jwtAuth, corsMid)
	e.POST("threadsave", cl.ThreadSaveController.SaveThread, jwtAuth, corsMid)
	e.POST("threadshare", cl.ThreadShareController.ShareThread, jwtAuth, corsMid)

	//additional
	e.POST("category", cl.CategoryController.CreateCategoryController, jwtAuth, corsMid)
	e.POST("reportcase", cl.ReportCaseController.CreateCaseController, jwtAuth, corsMid)
	e.POST("reputation", cl.ReputationController.CreateReputationController, jwtAuth, corsMid)
}
