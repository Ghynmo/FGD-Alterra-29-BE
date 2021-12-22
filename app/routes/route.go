package routes

import (
	"fgd-alterra-29/controllers/threads"
	userbadges "fgd-alterra-29/controllers/user_badges"
	"fgd-alterra-29/controllers/users"

	"github.com/labstack/echo/v4"
)

type ControllerList struct {
	UserController      users.UserController
	UserBadgeController userbadges.UserBadgeController
	ThreadController    threads.ThreadController
}

func (cl *ControllerList) RouteRegister(e echo.Echo) {
	e.GET("users", cl.UserController.GetUsersController)
	e.GET("profile/:id", cl.UserController.GetProfileController)
}
