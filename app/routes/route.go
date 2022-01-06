package routes

import (
	"fgd-alterra-29/controllers/categories"
	"fgd-alterra-29/controllers/comments"
	"fgd-alterra-29/controllers/follows"
	threadlikes "fgd-alterra-29/controllers/thread_likes"
	"fgd-alterra-29/controllers/threads"
	userbadges "fgd-alterra-29/controllers/user_badges"
	"fgd-alterra-29/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController       users.UserController
	UserBadgeController  userbadges.UserBadgeController
	ThreadController     threads.ThreadController
	CommentController    comments.CommentController
	FollowController     follows.FollowController
	CategoryController   categories.CategoryController
	ThreadLikeController threadlikes.ThreadLikeController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.GET("users", cl.UserController.GetUsersController)
	e.GET("profile/:id", cl.UserController.GetProfileController)
	e.GET("post/:id", cl.CommentController.GetProfileComments)
	e.GET("thread/:id", cl.ThreadController.GetProfileThreads)
	e.GET("followers/:id", cl.FollowController.GetFollowers)
	e.GET("following/:id", cl.FollowController.GetFollowing)
	e.GET("home/:id", cl.ThreadController.GetHomepageThreads)
	e.GET("recommendation/:id", cl.ThreadController.GetRecommendationThreads)
	e.GET("hotthread", cl.ThreadController.GetHotThreads)
	e.GET("search", cl.ThreadController.GetSearch)

	e.POST("comment", cl.CommentController.CreateCommentController)
	e.POST("threadlike", cl.ThreadLikeController.CreateLikes)
	e.DELETE("threadlike", cl.ThreadLikeController.DeleteLikes)

	e.POST("commentlike", cl.ThreadController.GetSearch)
	e.DELETE("commentunlike", cl.ThreadController.GetSearch)
	e.POST("threadsave", cl.ThreadController.GetSearch)
	e.POST("threadshare", cl.ThreadController.GetSearch)
}
