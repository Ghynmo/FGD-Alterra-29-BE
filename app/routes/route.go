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
	"fmt"
	"reflect"
	"strings"

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
	admin.Use(middleware.KeyAuth(func(key string, c echo.Context) (bool, error) {
		reqToken := c.Request().Header.Get("Authorization")
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]

		claims, _ := middlewares.ExtractClaims(reqToken)
		val := reflect.ValueOf(claims["Admin"])
		data := val.Bool()

		fmt.Println(data)
		return data, nil
	}))

	//Main Page of Admin
	admin.GET("", cl.UserController.GetDashboardController, jwtAuth)

	//List of Users (Admin Access)
	admin.GET("/users", cl.UserController.GetUsersController, jwtAuth)
	admin.GET("/users/search/:name", cl.UserController.GetUsersByName, jwtAuth)
	admin.PUT("/users/banned/:id", cl.UserController.BannedUser, jwtAuth)

	//List of Threads (Admin Access)
	admin.GET("/threads", cl.ThreadController.GetThreadsController, jwtAuth)
	admin.GET("/threads/search/:title", cl.ThreadController.GetThreadsByTitleController, jwtAuth)
	admin.PUT("/threads/thread/:id", cl.ThreadController.DeleteThread, jwtAuth)
	admin.PUT("/threads/thread/:id", cl.ThreadController.ActivateThread, jwtAuth)

	//List of Posts (Admin Access)
	admin.GET("/posts", cl.CommentController.GetPostsController, jwtAuth)
	admin.GET("/posts/search/:comment", cl.CommentController.GetPostsByCommentController, jwtAuth)
	admin.PUT("/posts/post/:id", cl.CommentController.UnactivatingPostController, jwtAuth)
	admin.PUT("/posts/post/:id", cl.CommentController.ActivatingPostController, jwtAuth)

	//List of Comment's Reports (Admin Access)
	// admin.GET("/comment-reports", cl.CommentReportController.AdminGetReports, jwtAuth)
	// admin.GET("/comment-reports/search/:category", cl.CommentReportController.GetReportsByCategoryController, jwtAuth)
	// admin.DELETE("/comment-reports/comment-report/:id", cl.CommentReportController.DeleteCommentReport, jwtAuth)

	//List of Thread's Reports (Admin Access)
	admin.GET("/thread-reports", cl.ThreadReportController.AdminGetReports, jwtAuth)
	admin.GET("/thread-reports/search/:category", cl.ThreadReportController.SearchReportsByCategoryController, jwtAuth)
	admin.PUT("/thread-reports/thread-report/:id", cl.ThreadReportController.SolvedThreadReport, jwtAuth)

	//Edit Profile page
	admin.GET("/edit/:id", cl.UserController.GetAdminSettingController, jwtAuth)
	admin.PUT("/edit/:id", cl.UserController.UpdateAdminProfile, jwtAuth)

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
	e.GET("profile/:id", cl.UserController.GetProfileController, jwtAuth)
	e.GET("post/:id", cl.CommentController.GetProfileCommentsController, jwtAuth)
	e.GET("thread/:id", cl.ThreadController.GetProfileThreads, jwtAuth)
	e.GET("followers", cl.FollowController.GetFollowers)
	e.GET("following", cl.FollowController.GetFollowing)
	e.GET("user/edit/:id", cl.UserController.GetUserSettingController, jwtAuth)
	e.PUT("user/edit/:id", cl.UserController.UpdateUserProfile, jwtAuth)

	e.GET("", cl.ThreadController.GetHotThreads)
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

	//Report Thread page
	e.GET("report-thread", cl.ReportCaseController.GetReportForm, jwtAuth)
	e.POST("report-thread", cl.ThreadReportController.CreateReportThread, jwtAuth)

}
