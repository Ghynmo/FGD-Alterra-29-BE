package routes

import (
	"fgd-alterra-29/controllers/categories"
	"fgd-alterra-29/controllers/comments"
	"fgd-alterra-29/controllers/follows"
	"fgd-alterra-29/controllers/threads"
	userbadges "fgd-alterra-29/controllers/user_badges"
	"fgd-alterra-29/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController      users.UserController
	UserBadgeController userbadges.UserBadgeController
	ThreadController    threads.ThreadController
	CommentController   comments.CommentController
	FollowController    follows.FollowController
	CategoryController  categories.CategoryController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.GET("users", cl.UserController.GetUsersController)
	e.GET("profile/:id", cl.UserController.GetProfileController)
	e.GET("post/:id", cl.CommentController.GetProfileComments)
	e.GET("thread/:id", cl.ThreadController.GetProfileThreads)
	e.GET("followers/:id", cl.FollowController.GetFollowers)
	e.GET("following/:id", cl.FollowController.GetFollowing)
}
