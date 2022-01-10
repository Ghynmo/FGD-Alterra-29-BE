package routes

import (
	"fgd-alterra-29/controllers/categories"
	commentlikes "fgd-alterra-29/controllers/comment_likes"
	"fgd-alterra-29/controllers/comments"
	"fgd-alterra-29/controllers/follows"
	threadlikes "fgd-alterra-29/controllers/thread_likes"
	threadsaves "fgd-alterra-29/controllers/thread_saves"
	threadshares "fgd-alterra-29/controllers/thread_shares"
	"fgd-alterra-29/controllers/threads"
	userbadges "fgd-alterra-29/controllers/user_badges"
	"fgd-alterra-29/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController        users.UserController
	UserBadgeController   userbadges.UserBadgeController
	ThreadController      threads.ThreadController
	CommentController     comments.CommentController
	FollowController      follows.FollowController
	CategoryController    categories.CategoryController
	ThreadLikeController  threadlikes.ThreadLikeController
	CommentLikeController commentlikes.CommentLikeController
	ThreadSaveController  threadsaves.ThreadSaveController
	ThreadShareController threadshares.ThreadShareController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.GET("home/:id", cl.ThreadController.GetHomepageThreads)
	e.GET("recommendation/:id", cl.ThreadController.GetRecommendationThreads)
	e.GET("hotthread", cl.ThreadController.GetHotThreads)
	e.GET("search", cl.ThreadController.GetSearch)

	e.GET("comment/reply/:id", cl.CommentController.GetReplyComments)
	e.POST("comment", cl.CommentController.CreateCommentController)
	e.POST("threadlike", cl.ThreadLikeController.CreateLikes)
	e.DELETE("threadlike", cl.ThreadLikeController.DeleteLikes)
	e.POST("commentlike", cl.CommentLikeController.CreateLikes)
	e.DELETE("commentlike", cl.CommentLikeController.DeleteLikes)
	e.POST("threadsave", cl.ThreadSaveController.SaveThread)
	e.DELETE("threadsave", cl.ThreadSaveController.Unsaves)
	e.POST("threadshare", cl.ThreadShareController.ShareThread)
}
